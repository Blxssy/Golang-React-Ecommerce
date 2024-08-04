package storage

import (
	"fmt"
	"log/slog"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}

type storage struct {
	db *gorm.DB
}

func NewStorage(logger *slog.Logger, config *config.Config) Storage {
	db, err := connectDatabase(config)
	if err != nil {
		logger.Error("Failure database connection")
	}
	logger.Info("Successfully connection to database")
	logger.Info("db", slog.String("port", config.Port))
	return &storage{
		db: db,
	}
}

func connectDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username,
		config.Database.Name, config.Database.Password)
	return gorm.Open(postgres.Open(dsn))

}

func (s *storage) Create(value interface{}) *gorm.DB {
	return s.db.Create(value)
}

func (s *storage) Find(out interface{}, where ...interface{}) *gorm.DB {
	return s.db.Find(out, where...)
}

func (s *storage) First(out interface{}, where ...interface{}) *gorm.DB {
	return s.db.First(out, where...)
}

func (s *storage) Where(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Where(query, args...)
}
