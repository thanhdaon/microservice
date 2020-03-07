package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func CrawlerWithProxy() {
	CrawlProxy()

	c := setupCrawlerWithProxy()

	for i := 0; i < 6; i++ {
		c.Visit("https://api.ipify.org/")
	}
}

func setupCrawlerWithProxy() *colly.Collector {
	c := colly.NewCollector(colly.AllowURLRevisit())

	c.SetProxy("http://200.89.174.158:3128")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("[ERROR]", err)
	})

	return c
}
