package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var proxies []string

func CrawlProxy() {
	proxies = []string{}

	c := setupCrawProxy()

	c.Visit("https://free-proxy-list.net/")
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
		https := e.ChildText("td:nth-child(7)")

		version := "http"
		if https == "yes" {
			version = "https"
		}
		proxy := fmt.Sprintf("%s://%s:%s\n", strings.ToLower(version), ip, port)
		if proxyOk(proxy) {
			fmt.Println("[OK] ", proxy)
			proxies = append(proxies, proxy)
		} else {
			fmt.Println("[FALSE] ", proxy)
		}

	})

	return c
}

func proxyOk(proxy string) bool {
	return false
	os.Setenv("HTTP_PROXY", proxy)
	client := &http.Client{}
	resp, err := client.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "google") {
		return true
	}
	return false
}
