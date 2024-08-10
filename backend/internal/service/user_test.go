package service

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_LoginUser(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewUserService(container)
	res, token, refresh, err := service.LoginUser("test@test.com", "test")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, res)
	assert.NotNil(t, refresh)
}

func setUpTestData(container container.Container) {
	entity := models.NewUserWithPlainPassword("test", "test@test.com", "test")
	repo := container.GetRepository()
	_, _ = entity.Create(repo)

	entity = models.NewUserWithPlainPassword("test2", "test2@test.com", "test")
	_, _ = entity.Create(repo)
}
