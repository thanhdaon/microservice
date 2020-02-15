package utils

import (
	"net/url"
)

func ExtractDomain(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
