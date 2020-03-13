package crawler

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

var re = regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+`)

func CrawlEmailFromSearchResult() {
	msgs, err := channel.Consume(
		BING_SEARCH_RESULT_QUEUE, // queue
		"",                       // consumer
		false,                    // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
			crawlEmailFromSearchResult(string(msg.Body))
			log.Printf("Proccessed a message: %s", msg.Body)
			msg.Ack(false)
		}
	}()

}

func crawlEmailFromSearchResult(url string) {
	if isResourceExist(url) {
		log.Println("Resource exist")
		return
	}

	domain, err := getDomain(url)
	if err != nil {
		log.Println("fail when get domain")
		return
	}

	if isJsRenderingWebsite(url) {
		publishToRabbit(url, JS_BASED_WEBSITE_QUEUE)
		log.Println("JS rendering website")
		return
	}

	addDomain(domain)

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowedDomains(domain),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		e.Request.Visit(link)
	})

	c.OnResponse(func(r *colly.Response) {
		emails := NewSet()

		for _, found := range re.FindAll(r.Body, 2) {
			if isValidEmail(string(found)) {
				emails.Add(string(found))
			}
		}

		var resource Resource
		db.First(&resource, "resource = ?", r.Request.URL.String())
		if resource.ID == 0 {
			resource = Resource{Resource: r.Request.URL.String()}
			db.Create(&resource)
		}
		for emailStr, _ := range emails.Iterator() {
			var email Email
			db.First(&email, "email = ?", emailStr)
			if email.ID == 0 {
				log.Printf("new email: %s\n", emailStr)
				email = Email{Email: emailStr}
				email.Resources = append(email.Resources, &resource)
				db.Create(&email)
			} else {
				if email.ResourceCount < 8 {
					email.ResourceCount = email.ResourceCount + 1
					db.Save(&email)
					db.Model(&email).Association("resources").Append(&resource)
				}
			}
		}
	})

	c.Visit(url)
}
