package controller

import (
	"net/http"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/avatar"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	service   service.UserService
	container container.Container
}

func NewUserController(container container.Container) UserController {
	return &userController{
		container: container,
		service:   service.NewUserService(container),
	}
}

func (u *userController) RegisterUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	result, _ := u.service.FindByEmail(input.Email)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	username := faker.Username()
	user := models.User{
		Email:      input.Email,
		Username:   username,
		PassHash:   string(hashedPassword),
		AvatarPath: avatar.GenerateRandomAvatar(username),
		Phone:      faker.Phonenumber(),
	}

	u.container.GetRepository().Create(&user)

	c.JSON(http.StatusOK, user)
}

func (u *userController) GetUsers(c *gin.Context) {

}

func (u *userController) GetUserByID(c *gin.Context) {

}

func (u *userController) CreateUser(c *gin.Context) {

}

func (u *userController) UpdateUser(c *gin.Context) {

}

func (u *userController) DeleteUser(c *gin.Context) {

}
