package migration

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
)

func CreateDatabase(container container.Container) {
	if container.GetConfig().Database.Migration {
		db := container.GetRepository()

		_ = db.DropTableIfExists(&models.User{})
		_ = db.DropTableIfExists(&models.Cart{})
		_ = db.DropTableIfExists(&models.CartItem{})
		_ = db.DropTableIfExists(&models.Product{})

		_ = db.AutoMigrate(&models.User{})
		_ = db.AutoMigrate(&models.Cart{})
		_ = db.AutoMigrate(&models.Category{})
		_ = db.AutoMigrate(&models.Product{})
		_ = db.AutoMigrate(&models.CartItem{})
	}
}
