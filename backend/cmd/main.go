package main

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"

	_ "github.com/Blxssy/Golang-React-Ecommerce/docs"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/logger"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/migration"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/router"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
)

func main() {
	exists := godotenv.Load(".env")
	if exists != nil {
		err := godotenv.Load("example.env")
		if err != nil {
			log.Println("Error loading .env file")
			os.Exit(1)
		}
	}

	token.InitJWTKey()

	g := gin.Default()

	g.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	cfg := config.NewConfig()

	l := logger.SetupLogger(envLocal)

	mainStorage := storage.NewStorage(l, cfg)

	ctr := container.NewContainer(mainStorage, cfg, l, envLocal)

	migration.CreateDatabase(ctr)
	migration.InitData(ctr)

	router.Init(g, ctr)

	g.Run(":3001")
}
