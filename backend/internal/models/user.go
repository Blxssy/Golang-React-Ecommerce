package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username   string `json:"username"`
	Email      string `json:"email"`
	PassHash   string `json:"pass_hash"`
	AvatarPath string `gotrm:"default:https://api.multiavatar.com/"`
	Phone      string `json:"phone"`
}
