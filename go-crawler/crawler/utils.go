package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func proxyOk(scheme, ip, port string) bool {
	host := fmt.Sprintf("%s:%s", ip, port)
	urlProxy := &url.URL{Scheme: scheme, Host: host}
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(urlProxy)},
		Timeout:   2 * time.Second,
	}
	resp, err := client.Get("https://market.m.taobao.com/app/tbhome/common/error.html")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(body), "taobao") {
		return true
	}
	return false
}

func getCompanyKeywords() []string {
	data, err := ioutil.ReadFile("company.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}
