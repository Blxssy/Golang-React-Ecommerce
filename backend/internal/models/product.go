package models

import (
	"fmt"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
}

type Category struct {
	gorm.Model

	Name     string
	Products []Product `gorm:"many2many:category_products;" json:"products"`
}

func (p *Product) Create(s storage.Storage) (*Product, error) {
	if err := s.Select("name", "price", "description", "slug", "image").Create(p).Error; err != nil {
		return nil, err
	}
	fmt.Println(p)
	return p, nil
}
