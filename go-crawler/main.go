package main

<<<<<<< Updated upstream
import "email-crawler/crawler"

func main() {
	crawler.CrawlProxy()
=======
import (
	"email-crawler/crawler"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

func main() {
	crawler.SetupRabbit()

	data, err := ioutil.ReadFile("log.txt")
	check(err)
	for _, url := range strings.Split(string(data), "\n") {
		crawler.Crawl(makeSearchUrl(url))
	}
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func makeSearchUrl(keyword string) string {
	params := url.Values{}
	params.Add("q", keyword)
	return fmt.Sprintf("http://www.google.com/search?%s&start=1", params.Encode())
}

func makeUrl(pageIndex int) string {
	return fmt.Sprintf("https://reviewcongty.com/?tab=latest&page=%d", pageIndex)
>>>>>>> Stashed changes
}
