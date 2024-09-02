package test

import (
	"encoding/json"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	logger2 "github.com/Blxssy/Golang-React-Ecommerce/internal/logger"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/migration"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
)

func PrepareForServiceTest() container.Container {
	conf := createConfig()
	logger := initTestLogger()
	container := initContainer(conf, logger)

	migration.CreateTestDB(container)
	migration.InitData(container)

	return container
}

func PrepareForControllerTest() (*gin.Engine, container.Container) {
	g := gin.Default()

	conf := createConfig()
	logger := initTestLogger()
	container := initContainer(conf, logger)

	migration.CreateTestDB(container)
	migration.InitData(container)

	return g, container
}

func createConfig() *config.Config {
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

func NewJSONRequest(method string, target string, param interface{}) *http.Request {
	req := httptest.NewRequest(method, target, strings.NewReader(ConvertToString(param)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

func ConvertToString(model interface{}) string {
	bytes, _ := json.Marshal(model)
	return string(bytes)
}
