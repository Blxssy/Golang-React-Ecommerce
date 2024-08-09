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
	g.POST(config.APIREFRESH, user.RefreshTokens)
	g.GET(config.APIUsers, user.GetUsers)
	g.GET(config.APIUsersID, user.GetUserByID)
	g.POST(config.APIUsers, user.CreateUser)
	g.PUT(config.APIUsersID, user.UpdateUser)
	g.DELETE(config.APIUsersID, user.DeleteUser)
}

func setProductController(g *gin.Engine, container container.Container) {
	product := controller.NewProducController(container)
	g.GET(config.APIProducts, product.GetProducts)
	g.POST(config.APIProducts, product.CreateProduct)
}
