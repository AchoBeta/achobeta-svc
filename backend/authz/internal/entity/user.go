package entity

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Nickname     string `json:"nickname"`
	Gender       int8   `json:"gender"` // 性别 0:未知 1:男 2:女
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	gorm.Model
}

func (*User) TableName() string {
	return "t_user"
}

func MockUser() *User {
	uid := uuid.NewString()
	return &User{
		Nickname:     fmt.Sprintf("用户%s", uid[len(uid)-10:]),
		Gender:       0,
		Avatar:       "",
		Introduction: "",
	}
}

type UserInfoEntity struct {
	Id       uint
	Username string
	Email    string
	Phone    string
	User     *User
}
