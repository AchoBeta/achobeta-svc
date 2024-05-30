package entity

import (
	"achobeta-svc/internal/achobeta-svc-common/pkg/utils"
	"fmt"

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
	return "ab_user"
}

func MockUser() *User {
	mockNickName := utils.GetSnowflakeUUID()[:6]
	return &User{
		Nickname:     fmt.Sprintf("用户%s", mockNickName),
		Gender:       0,
		Avatar:       "",
		Introduction: "",
	}
}
