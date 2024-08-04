package container

import (
	"log/slog"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
)

type Container interface {
	GetRepository() storage.Storage
	GetConfig() *config.Config
	GetLogger() *slog.Logger
	GetEnv() string
}

type container struct {
	storage storage.Storage
	config  *config.Config
	logger  *slog.Logger
	env     string
}

func NewContainer(
	storage storage.Storage,
	cfg *config.Config,
	log *slog.Logger,
	env string,
) *container {
	return &container{
		storage: storage,
		config:  cfg,
		logger:  log,
		env:     env,
	}
}

func (c *container) GetRepository() storage.Storage {
	return c.storage
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetLogger() *slog.Logger {
	return c.logger
}

func (c *container) GetEnv() string {
	return c.env
}
