package crawler

import (
	"fmt"
	"log"
	"time"

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

	c := colly.NewCollector(colly.AllowURLRevisit())

	proxies := CrawlProxy()
	if len(proxies) > 0 {
		rp, err := proxy.RoundRobinProxySwitcher(proxies...)
		if err != nil {
			log.Fatal(err)
		}
		c.SetProxyFunc(rp)
	}

	go func() {
		for {
			time.Sleep(30 * time.Second)
			proxies := CrawlProxy()
			if len(proxies) > 0 {
				rp, err := proxy.RoundRobinProxySwitcher(proxies...)
				if err != nil {
					log.Fatal(err)
				}
				c.SetProxyFunc(rp)
			}
		}
	}()

	c.SetRequestTimeout(4 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ON REQUEST]", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[SUCCESS]", string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("[ON ERROR] %s - proxy: %s \n", r.Request.URL, r.Request.ProxyURL)
		c.Visit(r.Request.URL.String())
	})

	return c
}
