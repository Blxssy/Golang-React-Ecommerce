package controller

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartController interface {
	AddItem(c *gin.Context)
	RemoveItem(c *gin.Context)
	UpdateItemQuantity(c *gin.Context)
	GetCart(c *gin.Context)
	ClearCart(c *gin.Context)
}

type cartController struct {
	container container.Container
	service   service.CartService
}

func NewCartController(container container.Container) CartController {
	return &cartController{
		container: container,
		service:   service.NewCartService(container),
	}
}

// GetCart take access_token from cookies and returns user's cart
// GetCart godoc
// @Summary Get user cart
// @Description Retrieves the shopping cart for a specific user using the access token stored in cookies.
// @Tags cart
// @Accept  json
// @Produce  json
// @Param access_token header string true "Access Token"
// @Success 200 {object} map[string]interface{} "{"uid": "string", "userCart": "object"}"
// @Failure 400 {object} map[string]interface{} "{"error": "string"}"
// @Router /api/cart [get]
func (cart *cartController) GetCart(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := token.ParseToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCart, err := cart.service.GetCart(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":      uid,
		"userCart": userCart,
	})
}

// AddItem godoc
// @Summary Add item to cart
// @Description Adds a product to the user's shopping cart using the access token stored in cookies.
// @Tags cart
// @Accept  json
// @Produce  json
// @Param access_token header string true "Access Token"
// @Param input body object{product_id=int,quantity=int} true "Product ID and Quantity"
// @Success 200 {object} map[string]interface{} "{"message": "Product successfully added to cart"}"
// @Failure 400 {object} map[string]interface{} "{"error": "string"}"
// @Router /api/cart/items [post]
func (cart *cartController) AddItem(c *gin.Context) {
	uid, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cart.service.AddItem(uid, input.ProductID, input.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully added to cart"})
}
func (cart *cartController) RemoveItem(c *gin.Context) {

}

func (cart *cartController) UpdateItemQuantity(c *gin.Context) {

}

func (cart *cartController) ClearCart(c *gin.Context) {

}
