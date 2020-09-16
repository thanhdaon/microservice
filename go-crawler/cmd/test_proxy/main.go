package main

import (
	"email-crawler/crawler"
	"log"
)

func main() {
	ok := crawler.ProxyOk("http", "115.74.201.137", "42108")
	log.Println(ok)
}
