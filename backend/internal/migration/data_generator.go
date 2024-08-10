package migration

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
)

func InitData(container container.Container) {
	rep := container.GetRepository()

	u := models.NewUserWithPlainPassword("test", "test@test.com", "test")
	u.Create(rep)

	u = models.NewUserWithPlainPassword("test2", "test2@test.com", "test2")
	u.Create(rep)
}
