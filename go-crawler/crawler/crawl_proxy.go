package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func CrawlProxy() []string {
	proxies := []string{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ON REQUEST]", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("[ON ERROR]", err)
	})

	c.OnHTML("#proxylisttable > tbody > tr", func(e *colly.HTMLElement) {
		ip := e.ChildText("td:nth-child(1)")
		port := e.ChildText("td:nth-child(2)")
		https := e.ChildText("td:nth-child(7)")
		scheme := "http"
		if strings.ToLower(https) == "yes" {
			scheme = "https"
		}
		proxy := fmt.Sprintf("%s://%s:%s", scheme, ip, port)
		proxies = append(proxies, proxy)
	})

	c.Visit("https://free-proxy-list.net/")

	fmt.Printf("[INFO] refresh proxies: %d new proxies\n", len(proxies))

	return proxies
}
