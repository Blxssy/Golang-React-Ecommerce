package service

import (
	"strconv"
	"testing"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestUserService_RegisterUser(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewUserService(container)
	res, token, refresh, err := service.RegisterUser("name", "register@test.com", "test")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, refresh)
}

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

func TestUserService_FindAllUsers(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewUserService(container)
	users, err := service.FindAllUsers()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestUserService_FindById(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)
	service := NewUserService(container)

	user, err := service.FindById(strconv.Itoa(1))

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUserService_FindByEmail(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)
	service := NewUserService(container)

	user, err := service.FindByEmail("test@test.com")

	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}

func TestUserService_CreateUser(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)
	service := NewUserService(container)

	user := models.NewUserWithPlainPassword("test3", "test3@test.com", "test")
	err := service.CreateUser(user)

	assert.NoError(t, err)
}

func setUpTestData(container container.Container) {
	entity := models.NewUserWithPlainPassword("test", "test@test.com", "test")
	repo := container.GetRepository()
	_, _ = entity.Create(repo)

	entity = models.NewUserWithPlainPassword("test2", "test2@test.com", "test")
	_, _ = entity.Create(repo)
}
