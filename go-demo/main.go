package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/imroc/req"
)

type ResponseBody struct {
	Posts []struct {
		ID    string `json:"id"`
		UUID  string `json:"uuid"`
		Title string `json:"title"`
	} `json:"posts"`
}

func main() {
	demo3()
}

func demo1() {
	resp, err := http.Get("https://trangnhat.net/tin-tuc/ghost/api/v3/content/posts?key=9f1e64bcb6124ed6d38debcc9b&limit=2&field=all")
	if err != nil {
		fmt.Println("1")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("2")
	}

	var response ResponseBody
	if json.Unmarshal(body, &response) != nil {
		fmt.Println("3")
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
