package main

import (
	"email-crawler/crawler"
	"fmt"
	"net/url"
)

func main() {
	crawler.Crawl(makeSearchUrl("okok"))
}

func makeSearchUrl(keywork string) string {
	params := url.Values{}
	params.Add("q", keywork)
	return fmt.Sprintf("http://www.google.com/search?%s&start=1", params.Encode())
}

func makeUrl(pageIndex int) string {
	return fmt.Sprintf("https://reviewcongty.com/?tab=latest&page=%d", pageIndex)
}
