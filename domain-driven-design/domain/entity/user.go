package entity

import "domain-driven-design/utils/auth"

type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func (u *User) BeforeSave() error {
	hashPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}
