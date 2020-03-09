package crawler

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func CrawlGoogle() {
	c := setupCrawlerWithProxy()

	for _, keyword := range getCompanyKeywords() {
		c.Visit(makeGoogleSearchUrl(keyword))
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

	c.SetRequestTimeout(4 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ON REQUEST]", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("[SUCCESS]", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("[ERROR] ", err)
		c.Visit(r.Request.URL.String())
	})

	c.OnHTML(".kCrYT a[href]", func(e *colly.HTMLElement) {
		publishToRabbit(extractLinkFromGoogleResult(e.Attr("href")))
	})

	return c
}

func makeGoogleSearchUrl(keyword string) string {
	params := url.Values{}
	params.Add("q", keyword)
	return fmt.Sprintf("http://www.google.com/search?%s&start=1", params.Encode())
}

func extractLinkFromGoogleResult(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		return ""
	}
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}
	return values.Get("q")
}
