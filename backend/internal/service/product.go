package service

import (
	"fmt"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
)

type ProductService interface {
	GetAllProducts() (*[]models.Product, error)
	CreateProduct(product *models.Product) error
}

type productService struct {
	container container.Container
}

func NewProductService(c container.Container) ProductService {
	return &productService{
		container: c,
	}
}

func (s *productService) GetAllProducts() (*[]models.Product, error) {
	var products []models.Product
	if err := s.container.GetRepository().Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (s *productService) CreateProduct(product *models.Product) error {
	rep := s.container.GetRepository()
	var err error
	fmt.Println(product)
	if txerr := rep.Transaction(func(tx storage.Storage) error {
		_, err = txCreateProduct(tx, product)
		return err
	}); txerr != nil {
		s.container.GetLogger().Error(txerr.Error())
		return nil
	}

	return nil
}

func txCreateProduct(txstorage storage.Storage, product *models.Product) (*models.Product, error) {
	var result *models.Product
	var err error

	if result, err = product.Create(txstorage); err != nil {
		return nil, err
	}

	return result, nil
}
