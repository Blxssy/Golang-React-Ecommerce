package models

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	UserID     uint       `json:"user_id"`
	Items      []CartItem `json:"items"`
	TotalPrice int        `json:"total_price"`
}

type CartItem struct {
	gorm.Model

	CartID    uint `json:"cart_id"`
	ProductId uint `json:"product_id" gorm:"foreignKey:ProductID"`
	Quantity  int  `json:"quantity"`
	Price     int  `json:"price"`
}

func (c *Cart) Save(s storage.Storage) error {

	if err := s.Save(c).Error; err != nil {
		return err
	}
	return nil
}
