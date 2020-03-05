package crawler

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly"
)

func Crawl(url string) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Printf("Execute time: %s\n", time.Since(start))
	}(start)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML(".kCrYT a[href]", func(e *colly.HTMLElement) {
		fmt.Println(extractLink(e.Attr("href")))
	})

	c.Visit(url)
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
