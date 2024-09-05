package controller

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController interface {
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
}

type orderController struct {
	service   service.OrderService
	container container.Container
}

func NewOrderController(container container.Container) OrderController {
	return &orderController{
		container: container,
		service:   service.NewOrderService(container),
	}
}

func (oc *orderController) CreateOrder(c *gin.Context) {
	uid, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.service.CreateOrder(uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Ok": "Order created"})
}

func (oc *orderController) GetOrders(c *gin.Context) {
	uid, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := oc.service.GetOrders(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot get orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Orders": orders})
}
