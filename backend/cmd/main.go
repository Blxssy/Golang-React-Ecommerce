package main

import (
	"log"
	"os"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
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
