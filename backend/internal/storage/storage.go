package storage

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	Model(value interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Updates(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB
	ScanRows(rows *sql.Rows, result interface{}) error
	Transaction(fc func(tx Storage) error) (err error)
	Close() error
	DropTableIfExists(value interface{}) error
	AutoMigrate(value interface{}) error
}

type storage struct {
	db *gorm.DB
}

func NewStorage(logger *slog.Logger, config *config.Config) Storage {
	db, err := connectDatabase(config)
	if err != nil {
		logger.Error("Failure database connection")
	}
	//logger.Info("Successfully connection to database")
	//logger.Info("db", slog.String("port", config.Port))
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

func (s *storage) Model(value interface{}) *gorm.DB {
	return s.db.Model(value)
}

func (s *storage) Select(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Select(query, args...)
}

func (s *storage) Find(out interface{}, where ...interface{}) *gorm.DB {
	return s.db.Find(out, where...)
}

func (s *storage) Exec(sql string, values ...interface{}) *gorm.DB {
	return s.db.Exec(sql, values...)
}

func (s *storage) First(out interface{}, where ...interface{}) *gorm.DB {
	return s.db.First(out, where...)
}

func (s *storage) Raw(sql string, values ...interface{}) *gorm.DB {
	return s.db.Raw(sql, values...)
}

func (s *storage) Create(value interface{}) *gorm.DB {
	return s.db.Create(value)
}

func (s *storage) Save(value interface{}) *gorm.DB {
	return s.db.Save(value)
}

func (s *storage) Updates(value interface{}) *gorm.DB {
	return s.db.Updates(value)
}

func (s *storage) Delete(value interface{}) *gorm.DB {
	return s.db.Delete(value)
}

func (s *storage) Where(query interface{}, args ...interface{}) *gorm.DB {
	return s.db.Where(query, args...)
}

func (s *storage) Preload(column string, conditions ...interface{}) *gorm.DB {
	return s.db.Preload(column, conditions...)
}

func (s *storage) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return s.db.Scopes(funcs...)
}

func (s *storage) ScanRows(rows *sql.Rows, result interface{}) error {
	return s.db.ScanRows(rows, result)
}

func (s *storage) Close() error {
	sqlDB, _ := s.db.DB()
	return sqlDB.Close()
}

func (s *storage) DropTableIfExists(value interface{}) error {
	return s.db.Migrator().DropTable(value)
}

func (s *storage) AutoMigrate(value interface{}) error {
	return s.db.AutoMigrate(value)
}

func (s *storage) Transaction(fc func(tx Storage) error) (err error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Отмена транзакции в случае ошибки или паники
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // Повторное выбрасывание паники после отката
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	// Выполнение функции транзакции
	if err = fc(&storage{db: tx}); err != nil {
		return err
	}

	return nil
}
