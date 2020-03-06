package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

func CrawlProxy() {
	c := setupCrawProxy()

	c.Visit("https://www.socks-proxy.net/")
}

func setupCrawProxy() *colly.Collector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("[ERROR]", err)
	})

	c.OnHTML("tr", func(r *colly.HTMLElement) {
		r.ForEach("td", func(i int, r *colly.HTMLElement) {
			fmt.Print(r.Text, "-")
		})
		fmt.Println("")
	})

	return c
}
