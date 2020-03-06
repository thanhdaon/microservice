package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func CrawlEmail() {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Printf("Execute time: %s\n", time.Since(start))
	}(start)

	c := colly.NewCollector()

	rp, err := proxy.RoundRobinProxySwitcher(
		"https://www.us-proxy.org/",
		"https://free-proxy-list.net/anonymous-proxy.html",
		"https://www.socks-proxy.net/",
		"https://www.sslproxies.org/",
	)

	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML(".kCrYT a[href]", func(e *colly.HTMLElement) {
<<<<<<< Updated upstream:go-crawler/crawler/crawler_email.go
		publishToRabbit(extractLink(e.Attr("href")))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
=======
		url := extractLink(e.Attr("href"))
		publishToRabbit(url)
>>>>>>> Stashed changes:go-crawler/crawler/crawler.go
	})

	data, err := ioutil.ReadFile("log.txt")
	check(err)
	for _, keyword := range strings.Split(string(data), "\n") {
		c.Visit(makeSearchUrl(keyword))
		break
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func makeSearchUrl(keyword string) string {
	params := url.Values{}
	params.Add("q", keyword)
	return fmt.Sprintf("http://www.google.com/search?%s&start=1", params.Encode())
}

func makeUrl(pageIndex int) string {
	return fmt.Sprintf("https://reviewcongty.com/?tab=latest&page=%d", pageIndex)
}

func extractLink(urlString string) string {
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
