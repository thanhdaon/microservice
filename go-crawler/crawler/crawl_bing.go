package crawler

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func CrawlBing() {
	c := setupBingCrawler()

	for _, keyword := range getCompanyKeywords() {
		c.Visit(makeBingSearchUrl(keyword))
	}

	// c.Visit(makeBingSearchUrl("fpt"))
}

func setupBingCrawler() *colly.Collector {
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
			time.Sleep(5 * time.Minute)
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ON REQUEST]", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("[ERROR] ", r.Request.URL)
		c.Visit(r.Request.URL.String())
	})

	c.OnHTML(".b_algo h2 a[href]", func(e *colly.HTMLElement) {
		fmt.Println("[SUCCESS] ", e.Attr("href"))
		publishToRabbit(e.Attr("href"))
	})

	return c
}

func makeBingSearchUrl(query string) string {
	params := url.Values{}
	params.Add("q", query)
	return fmt.Sprintf("https://www.bing.com/search?%s&t=h_&ia=web", params.Encode())
}
