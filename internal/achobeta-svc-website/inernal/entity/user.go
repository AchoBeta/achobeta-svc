package entity

import "gorm.io/gorm"

type User struct {
	Name         string `json:"name"`
	Gender       int8   `json:"gender"` // 性别 0:未知 1:男 2:女
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	gorm.Model
}

func (*User) TableName() string {
	return "ab_user"
}
