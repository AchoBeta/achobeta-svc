package user

import "gorm.io/gorm"

type UserEntity struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
	gorm.Model
}
