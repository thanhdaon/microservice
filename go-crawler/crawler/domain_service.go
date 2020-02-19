package crawler

func isValidDomain(domainStr string) bool {
	var domain Domain
	db.Where("domain = ?", domainStr).First(&domain)
	if domain.ID == 0 {
		return true
	}
	return false
}

func addDomain(domainStr string) {
	db.Create(&Domain{Domain: domainStr})
}
