package service

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
)

type OrderService interface {
	CreateOrder(uid uint) error
	GetOrders(uid uint) ([]models.Order, error)
}

type orderService struct {
	container container.Container
}

func NewOrderService(container container.Container) OrderService {
	return &orderService{
		container: container,
	}
}

func (o *orderService) CreateOrder(uid uint) error {
	var cart models.Cart
	if err := o.container.GetRepository().Preload("Items").Where("user_id = ?", uid).First(&cart).Error; err != nil {
		return err
	}

	order := models.Order{
		UserID:     uid,
		Status:     "pending",
		TotalPrice: float64(cart.TotalPrice),
	}

	if err := o.container.GetRepository().Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrders(uid uint) ([]models.Order, error) {
	var orders []models.Order
	if err := o.container.GetRepository().Where("user_id = ?", uid).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
