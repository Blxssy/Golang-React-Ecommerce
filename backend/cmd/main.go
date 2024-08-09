package main

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"log"
	"os"

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

	config := config.NewConfig()

	logger := logger.SetupLogger(envLocal)

	storage := storage.NewStorage(logger, config)

	container := container.NewContainer(storage, config, logger, envLocal)

	migration.CreateDatabase(container)

	router.Init(g, container)

	g.Run(":3001")
}
