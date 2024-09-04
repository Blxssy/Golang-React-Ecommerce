package router

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/controller"
	token2 "github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, container container.Container) {
	setCORSConfig(g)
	setUserController(g, container)
	setProductController(g, container)
	setCartController(g, container)
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		uid, err := token2.VerifyToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", uid)
		c.Next()
	}
}
