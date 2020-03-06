package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

var proxies []string

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

	c.OnHTML("#proxylisttable > tbody > tr", func(e *colly.HTMLElement) {
		ip := e.ChildText("td:nth-child(1)")
		port := e.ChildText("td:nth-child(2)")
		version := e.ChildText("td:nth-child(5)")

		if proxyOk(ip, port) {
			proxies = append(proxies, fmt.Sprintf("%s://%s:%s\n", strings.ToLower(version), ip, port))
		}

	})

	return c
}

func proxyOk(ip, port string) bool {
	host := fmt.Sprintf("%s:%s", ip, port)
	url_proxy := &url.URL{Host: host}
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(url_proxy)},
	}
	resp, err := client.Get("http://err.taobao.com/error1.html")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "alibaba.com") {
		return true
	}
	return false
}
