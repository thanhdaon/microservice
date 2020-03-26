package crawler

import (
	"log"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

func CrawlHaymora() {
	c := setupHaymoraCrawler()
	c.Visit("www.lacviet.com.vn")
}

func setupHaymoraCrawler() *colly.Collector {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Printf("[INFO] visiting %s\n", r.URL)
	})

	c.OnHTML(".infor__compnay a[rel=nofollow]", func(e *colly.HTMLElement) {
		url := e.DOM.Text()
		log.Printf("[INFO] url: %s\n", url)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(string(r.Body))
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("[ERROR] when visting %s: %+v", r.Request.URL, err)
		if redurectURL, err := getRedirect(r.Request.URL.String()); err == nil {
			log.Printf("[INFO] redirect to %s\n", redurectURL)
		}
	})

	return c
}

func getRedirect(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "can not get")
	}

	return resp.Request.URL.String(), nil
}
