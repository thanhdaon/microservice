package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

	for _, post := range response.Posts {
		fmt.Println(post)
	}
}

func demo2() {
	query := make(url.Values)
	query.Add("foo3", "123")
	url := &url.URL{RawQuery: query.Encode(), Host: "foo", Scheme: "https"}
	fmt.Println(url)
}

func demo3() {
	req := req.New()
	resp, err := req.Get("https://trangnhat.net/tin-tuc/ghost/api/v3/content/posts?key=9f1e64bcb6124ed6d38debcc9b&limit=2&field=all")
	if err != nil {
		fmt.Println("1")
	}
	var body ResponseBody
	if err := resp.ToJSON(&body); err != nil {
		fmt.Println("2")
	}
	for _, post := range body.Posts {
		fmt.Println(post.ID)
	}

}
