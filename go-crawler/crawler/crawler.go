package crawler

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

func crawl(url string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Printf("Execute time: %s\n", time.Since(start))
	}(start)

	re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	allowedDomain := extractDomain(url)

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowedDomains(allowedDomain),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if isUrl(link) && extractDomain(link) != allowedDomain {
			publishToRabbit(link)
		}
		e.Request.Visit(link)
	})

	c.OnResponse(func(r *colly.Response) {
		emails := NewSet()

		for _, emailStr := range re.FindAll(r.Body, 2) {
			emails.Add(string(emailStr))
		}

		var resource Resource
		db.First(&resource, "resource = ?", r.Request.URL.String())
		if resource.ID == 0 {
			resource = Resource{Resource: r.Request.URL.String()}
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

func extractDomain(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	return u.Hostname()
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
