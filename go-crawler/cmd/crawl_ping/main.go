package main

import (
	"email-crawler/crawler"
)

func main() {
	crawler.SetupRabbit()
	crawler.CrawlBing()
}
