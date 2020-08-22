package crawler

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func CrawlFacebook() {
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("[ERROR] %v\n", err)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", os.Getenv("COOKIE"))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.Visit("https://mbasic.facebook.com/")
}
