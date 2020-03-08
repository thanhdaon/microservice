package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

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

	fmt.Println("[INFO] refresh proxies")

	return proxies
}

func proxyOk(scheme, ip, port string) bool {
	host := fmt.Sprintf("%s:%s", ip, port)
	urlProxy := &url.URL{Scheme: scheme, Host: host}
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(urlProxy)},
		Timeout:   2 * time.Second,
	}
	resp, err := client.Get("https://market.m.taobao.com/app/tbhome/common/error.html")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "taobao") {
		return true
	}
	return false
}
