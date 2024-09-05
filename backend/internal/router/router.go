package router

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, container container.Container) {
	setCORSConfig(g)
	setUserController(g, container)
	setProductController(g, container)
	setCartController(g, container)
	setOrderController(g, container)
}

func setCORSConfig(g *gin.Engine) {
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
}

func setUserController(g *gin.Engine, container container.Container) {
	user := controller.NewUserController(container)
	g.POST(config.APIREGISTER, user.RegisterUser)
	g.POST(config.APILOGIN, user.Login)

	g.GET(config.APIUsers, user.GetUsers)
	g.GET(config.APIUsersID, user.GetUserByID)
	g.POST(config.APIUsers, user.CreateUser)
	g.PUT(config.APIUsersID, user.UpdateUser)
	g.DELETE(config.APIUsersID, user.DeleteUser)

	protected := g.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.GET(config.APIUSERINFO, user.GetInfo)
		protected.POST(config.APIREFRESH, user.RefreshTokens)
	}
}

func setProductController(g *gin.Engine, container container.Container) {
	product := controller.NewProductController(container)
	g.GET(config.APIProducts, product.GetProducts)
	g.POST(config.APIProducts, product.CreateProduct)
}

func setCartController(g *gin.Engine, container container.Container) {
	cart := controller.NewCartController(container)

	protected := g.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST(config.APICARTITEMS, cart.AddItem)
		protected.DELETE(config.APICARTITEMS, cart.RemoveItem)
		protected.PUT(config.APICARTITEMS, cart.UpdateItemQuantity)
		protected.GET(config.APICART, cart.GetCart)
		protected.DELETE(config.APICART, cart.ClearCart)
	}
}

func setOrderController(g *gin.Engine, container container.Container) {
	order := controller.NewOrderController(container)

	protected := g.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST(config.APICREATEORDER, order.CreateOrder)
		protected.GET(config.APIORDER, order.GetOrders)
	}
}
