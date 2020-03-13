package main

import (
	"email-crawler/crawler"
)

func main() {
	crawler.SetupDB()
	crawler.SetupRabbit()
	defer crawler.CleanupDB()
	defer crawler.CleanupRabbit()

	forever := make(chan bool)
	crawler.CrawlEmailFromSearchResult()
	crawler.CrawlEmailFromSearchResult()
	crawler.CrawlEmailFromSearchResult()
	<-forever
}
