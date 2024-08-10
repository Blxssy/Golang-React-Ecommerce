package models

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username   string `json:"username"`
	Email      string `json:"email" gorm:"unique"`
	PassHash   string `json:"pass_hash"`
	AvatarPath string `gorm:"default:https://api.multiavatar.com/"`
	Phone      string `json:"phone"`
}

// func NewUser() *User{
// 	return &User{

// 	}
// }

func (u *User) Create(s storage.Storage) (*User, error) {
	if err := s.Select("username", "pass_hash", "avatar_path", "email", "phone").Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func NewUserWithPlainPassword(name string, email string, password string) *User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Model:    gorm.Model{},
		Username: name,
		Email:    email,
		PassHash: string(hashed),
		Phone:    faker.Phonenumber(),
	}
}
