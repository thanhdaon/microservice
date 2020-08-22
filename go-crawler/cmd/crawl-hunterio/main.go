package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HunterSearchResult struct {
	Data struct {
		Domain string
		Emails []struct {
			Value   string
			Sources []struct {
				Domain string
				Uri    string
			}
		}
	}
}

func main() {
	fmt.Println("Start")
	file, err := os.Open("static/hunter-results.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	var results []HunterSearchResult
	if err = json.Unmarshal(byteValue, &results); err != nil {
		log.Fatalln(err)
	}

	for _, r := range results {
		fmt.Println("-------------------------")
		fmt.Printf("[domain] %s\n", r.Data.Domain)
		for _, e := range r.Data.Emails {
			fmt.Printf("[email] %s\n", e.Value)
			for _, s := range e.Sources {
				saveEmail(e.Value, s.Uri)
			}
		}
	}
}

func getDomain(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(u.Host, "www.", ""), nil
}

func domainsInHunter(domain string) ([]string, error) {
	url := fmt.Sprintf("https://api.congty.io/v1/emails/domain?search=%s", domain)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []struct {
		ID   string `json:"ID"`
		Text string `json:"Text"`
	}

	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	ret := []string{}
	for _, r := range results {
		ret = append(ret, r.Text)
	}

	return ret, nil
}

func fetchEmail(domain string) (interface{}, error) {
	url := fmt.Sprintf(
		"https://api.hunter.io/v2/domain-search?domain=%s&api_key=9b69408de270bded67b4fb75424dd1f6e5005188",
		domain,
	)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ret interface{}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func saveEmail(email, source string) error {
	client := &http.Client{}

	json, err := json.Marshal(map[string]string{
		"email":    email,
		"resource": source,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, "http://congty-io_email_service:8000/email", bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("[SaveEmail] %d\n", resp.StatusCode)
	return nil
}

func writeToFile(path string, data []interface{}) {
	file, _ := json.Marshal(data)
	_ = ioutil.WriteFile(path, file, 0644)
}

func readFromFile(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	var results []HunterSearchResult
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		log.Fatalln(err)
	}

	for _, r := range results {
		fmt.Println(r.Data.Domain)
	}
}
