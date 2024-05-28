package entity

import "gorm.io/gorm"

type User struct {
	Username string
	Password string
	Phone    string
	Email    string `gorm:"not null"`
	gorm.Model
}

func (*User) TableName() string {
	return "ab_user"
}
