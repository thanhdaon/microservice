package crawler

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
)

func CrawlDuckDuckGo() {
	c := setupDuckDuckGoCrawler()

	// for _, keyword := range getCompanyKeywords() {
	// 	c.Visit(makeDuckDuckGoSearchUrl(keyword))
	// }

	c.Visit(makeDuckDuckGoSearchUrl("fpt"))

}

func setupDuckDuckGoCrawler() *colly.Collector {
	c := colly.NewCollector(colly.AllowURLRevisit())

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ON REQUEST]", r.URL)
	})

	c.OnHTML(".b_algo h2 a[href]", func(e *colly.HTMLElement) {
		fmt.Println("[INFo] ", e.Attr("href"))
	})

	return c
}

func makeDuckDuckGoSearchUrl(query string) string {
	params := url.Values{}
	params.Add("q", query)
	return fmt.Sprintf("https://duckduckgo.com/?%s&t=h_&ia=web", params.Encode())
}
