package main

import (
	"email-crawler/crawler"
)

func main() {
	crawler.SetupDB()
	crawler.SetupRabbit()
	defer crawler.CleanupDB()
	defer crawler.CleanupRabbit()

	crawler.CrawlEmailFromSearchResult()
}
