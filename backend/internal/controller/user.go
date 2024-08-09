package controller

import (
	"errors"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"net/http"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	RefreshTokens(c *gin.Context)

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

	user, accessToken, refreshToken, err := u.service.RegisterUser(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.SetCookie("access_token", accessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "localhost", false, true)

	c.JSON(http.StatusOK, response)
}

func (u *userController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	usr, accessToken, refreshToken, err := u.service.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", accessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, 604800, "/", "localhost", false, true)

	response := map[string]interface{}{
		"user": gin.H{
			"id":    usr.ID,
			"email": usr.Email,
		},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusOK, response)
}

func (u *userController) RefreshTokens(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := token.UpdateToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("accessToken", accessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refreshToken", refreshToken, 604800, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.service.FindAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot load users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *userController) GetUserByID(c *gin.Context) {

}

func (u *userController) CreateUser(c *gin.Context) {

}

func (u *userController) UpdateUser(c *gin.Context) {

}

func (u *userController) DeleteUser(c *gin.Context) {

}
