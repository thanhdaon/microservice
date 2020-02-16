package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"email-crawler/pkg/database"
	"email-crawler/pkg/services"
	"email-crawler/pkg/utils"

	"github.com/gocolly/colly"
)

func main() {
	database.Setup()
	services.Setup()
	defer database.Close()
	crawl("https://dantri.com.vn/")
}

func crawl(url string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Printf("Execute time: %s\n", time.Since(start))
	}(start)

	re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	allowedDomain := utils.ExtractDomain(url)

	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowedDomains(allowedDomain),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnResponse(func(r *colly.Response) {
		emails := utils.NewSet()

		for _, emailStr := range re.FindAll(r.Body, 2) {
			emails.Add(string(emailStr))
		}

		var resource database.Resource
		database.DB.First(&resource, "resource = ?", r.Request.URL.String())
		if resource.ID == 0 {
			resource = database.Resource{Resource: r.Request.URL.String()}
		}
		for emailStr, _ := range emails.Iterator() {
			var email database.Email
			database.DB.First(&email, "email = ?", emailStr)
			if email.ID == 0 {
				log.Printf("new email: %s\n", emailStr)
				email = database.Email{Email: emailStr}
				email.Resources = append(email.Resources, &resource)
				database.DB.Create(&email)
			} else {
				if email.ResourceCount < 8 {
					email.ResourceCount = email.ResourceCount + 1
					database.DB.Save(&email)
					database.DB.Model(&email).Association("resources").Append(&resource)
				}
			}
		}
	})

	c.Visit(url)
}
