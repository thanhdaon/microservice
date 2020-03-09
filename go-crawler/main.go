package main

import (
	"email-crawler/crawler"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Printf("Execute time: %s\n", time.Since(start))
	}(start)

	crawler.SetupRabbit()
	// crawler.CrawlGoogle()
	// crawler.CrawlDuckDuckGo()
	crawler.CrawlBing()
}
