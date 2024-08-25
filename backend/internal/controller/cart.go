package controller

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/gin-gonic/gin"
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

func (cart *cartController) AddItem(c *gin.Context) {
	//userID, err := token.ParseToken(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}
	//
	//log.Println(userID)
	//
	//var req struct {
	//	ProductID uint `json:"product_id"`
	//	Quantity  int  `json:"quantity"`
	//}
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}
	//
	//err = cart.service.AddItem(userID, req.ProductID, req.Quantity)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}
func (cart *cartController) RemoveItem(c *gin.Context) {

}
func (cart *cartController) UpdateItemQuantity(c *gin.Context) {

}
func (cart *cartController) GetCart(c *gin.Context) {
	//userID, err := token.ParseToken(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}
	//
	//userCart, err := cart.service.GetCart(userID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart"})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, userCart)
}
func (cart *cartController) ClearCart(c *gin.Context) {

}
