package crawler

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func CrawlSerpapi() {
	c := setupSerpapiCrawler()

	for _, keyword := range getCompanyKeywords() {
		c.Visit(makeSerpapiSearchUrl(keyword))
	}
}

func setupSerpapiCrawler() *colly.Collector {
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
			time.Sleep(60 * time.Minute)
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

	c.OnResponse(func (r *colly.Response) {
		fmt.Println("[SUCCESS]", string(r.Body))
	})

	return c
}

func makeSerpapiSearchUrl(query string) string {
	params := url.Values{}
	params.Add("q", query)
	params.Add("location", "Vietnam")
	params.Add("hl", "vi")
	params.Add("gl", "vi")
	params.Add("google_domain", "google.com")
	return fmt.Sprintf("https://serpapi.com/search.json?%s", params.Encode())
}
