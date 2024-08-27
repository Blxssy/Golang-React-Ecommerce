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
	Email      string `json:"email"`
	PassHash   string `json:"pass_hash"`
	AvatarPath string `gorm:"default:https://api.multiavatar.com/"`
	Phone      string `json:"phone"`

	CartID uint `json:"cart_id"`
	//Cart   Cart `json:"cart"`
}

func (u *User) Create(s storage.Storage) (*User, error) {
	if err := s.Select("username", "pass_hash", "avatar_path", "email", "phone", "cart_id").Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	cart := Cart{UserID: u.ID}
	if err := tx.Create(&cart).Error; err != nil {
		return err
	}

	u.CartID = u.ID

	if err := tx.Save(u).Error; err != nil {
		return err
	}
	return nil
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
