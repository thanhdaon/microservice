package main

import (
	"fmt"
	"regexp"
	"time"

	"email-crawler/pkg/db"
	"email-crawler/pkg/set"
	"email-crawler/pkg/utils"

	"github.com/gocolly/colly"
)

func main() {
	db.Setup()
	defer db.Close()
	// crawl("https://caferati.me/")
}

func crawl(url string) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Execute time: %s", elapsed)
	}()

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
		emails := set.NewSet()

		for _, email := range re.FindAll(r.Body, 2) {
			emails.Add(string(email))
		}

		for email, _ := range emails.Iterator() {
			fmt.Println(email)
		}
	})

	c.Visit(url)
}
