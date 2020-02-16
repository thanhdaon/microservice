package database

type Domain struct {
	ID     int
	Domain string
}

type Email struct {
	ID            int
	Email         string
	ResourceCount int         `gorm:"default:1"`
	Resources     []*Resource `gorm:"many2many:email_resource"`
}

type Resource struct {
	ID       int
	Resource string
	Emails   []*Email `gorm:"many2many:email_resource"`
}
