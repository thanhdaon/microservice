package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func CrawlerWithProxy() {
	c := setupCrawlerWithProxy()

	for i := 0; i < 6; i++ {
		c.Visit("https://api.ipify.org/")
	}
}

func setupCrawlerWithProxy() *colly.Collector {
	CrawlProxy()

	c := colly.NewCollector(colly.AllowURLRevisit())

	if len(proxies) > 0 {
		rp, err := proxy.RoundRobinProxySwitcher(proxies...)
		if err != nil {
			log.Fatal(err)
		}
		c.SetProxyFunc(rp)
	}

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
