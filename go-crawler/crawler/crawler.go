package crawler

import (
	"fmt"
	"strconv"
	"strings"
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

	c.OnHTML(".pagination-list .pagination-link", func(e *colly.HTMLElement) {
		index, err := strconv.Atoi(e.Text)
		if err == nil {
			e.Request.Visit(makeUrl(index))
		}
	})

	c.OnHTML(".company-info__detail a[href]", func(e *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(e.Text))
	})

	c.Visit(url)
}

func makeUrl(pageIndex int) string {
	return fmt.Sprintf("https://reviewcongty.com/?tab=latest&page=%d", pageIndex)
}
