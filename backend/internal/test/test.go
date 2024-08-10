package test

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	logger2 "github.com/Blxssy/Golang-React-Ecommerce/internal/logger"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/migration"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"log/slog"
)

func PrepareForServiceTest() container.Container {
	conf := createConfig(false)
	logger := initTestLogger()
	container := initContainer(conf, logger)

	migration.CreateDatabase(container)
	migration.InitData(container)

	return container
}

func createConfig(isSecurity bool) *config.Config {
	conf := &config.Config{}
	conf.Database.Dialect = "postgres"
	conf.Database.Host = "localhost"
	conf.Database.Name = "test"
	conf.Database.Port = "5432"
	conf.Database.Username = "postgres"
	conf.Database.Password = "postgres"
	conf.Database.Migration = true
	return conf
}

func initContainer(conf *config.Config, logger *slog.Logger) container.Container {
	rep := storage.NewStorage(logger, conf)
	container := container.NewContainer(rep, conf, logger, "test")
	return container
}

func initTestLogger() *slog.Logger {
	logger := logger2.SetupLogger("local")
	return logger
}
