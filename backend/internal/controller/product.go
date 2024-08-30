package controller

import (
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

func NewProductController(c container.Container) ProductController {
	return &productController{
		container: c,
		service:   service.NewProductService(c),
	}
}

// GetProducts godoc
// @Summary Get all products
// @Description Retrieves a list of all available products.
// @Tags Products
// @Produce  json
// @Success 200 {array} object{ID=uint,name=string,price=int,description=string,slug=string,image=string} "List of products retrieved successfully"
// @Failure 400 {object} map[string]string "Products not found"
// @Router /api/products [get]
func (pc *productController) GetProducts(c *gin.Context) {
	products, err := pc.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "products not found"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Creates a new product with the provided information.
// @Tags Products
// @Accept  json
// @Produce  json
// @Param   product   body    object{name=string,price=int,description=string,slug=string,image=string}  true  "Product data"
// @Success 200 {object} object{ID=uint,name=string,price=int,description=string,slug=string,image=string} "Product created successfully"
// @Failure 400 {object} map[string]string "Invalid product data / Cannot create product"
// @Router /api/products [post]
func (pc *productController) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product data"})
		return
	}

	if err := pc.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
