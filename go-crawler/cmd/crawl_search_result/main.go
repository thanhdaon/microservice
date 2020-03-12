package main

import (
	"email-crawler/crawler"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://ahfarmer.github.io/emoji-search/")
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(crawler.IsJsRenderingWebsite(string(body)))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
