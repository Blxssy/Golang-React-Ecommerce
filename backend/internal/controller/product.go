package controller

import (
	"log/slog"
	"net/http"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetProducts(c *gin.Context)
	// GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
	// UpdateProduct(c *gin.Context)
	// DeleteProduct(c *gin.Context)
}

type productController struct {
	service   service.ProductService
	container container.Container
}

func NewProducController(c container.Container) ProductController {
	return &productController{
		container: c,
		service:   service.NewProductService(c),
	}
}

func (pc *productController) GetProducts(c *gin.Context) {
	products, err := pc.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "products not found"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *productController) CreateProduct(c *gin.Context) {
	var product models.Product
	logger := pc.container.GetLogger()
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}
	logger.Info("productS", slog.Any("product", product))
	if err := pc.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create product"})
		return
	}
	c.JSON(http.StatusOK, product)
}
