package services

import (
	"email-crawler/pkg/database"
)

func IsValidDomain(domainStr string) bool {
	var domain database.Domain
	db.Where("domain = ?", domainStr).First(&domain)
	if domain.ID == 0 {
		return true
	}
	return false
}

func AddDomain(domainStr string) {
	db.Create(&database.Domain{Domain: domainStr})
}
