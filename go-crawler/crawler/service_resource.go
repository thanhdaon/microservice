package crawler

func isResourceExist(target string) bool {
	var resource Resource
	db.Where("resource = ?", target).First(&resource)
	return resource.ID != 0
}
