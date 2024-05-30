package entity

import "gorm.io/gorm"

type Account struct {
	Username string `gorm:"not null"`
	UserId   uint   `gorm:"not null"`
	Password string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Email    string `gorm:"not null"`
	gorm.Model
}

func (*Account) TableName() string {
	return "ab_account"
}
