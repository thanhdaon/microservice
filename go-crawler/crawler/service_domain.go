package crawler

import "net/url"

func addDomain(domainStr string) {
	var domain Domain
	db.Where("domain = ?", domainStr).First(&domain)
	if domain.ID == 0 {
		db.Create(&Domain{Domain: domainStr})
	}
}

func extractDomain(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
