package models

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	UserID uint       `json:"user_id"`
	Items  []CartItem `json:"items"`
}

type CartItem struct {
	gorm.Model

	CartID    uint `json:"cart_id"`
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (c *Cart) Save(s storage.Storage) error {

	if err := s.Save(c).Error; err != nil {
		return err
	}
	return nil
}
