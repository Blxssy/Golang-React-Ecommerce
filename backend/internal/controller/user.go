package controller

import (
	"errors"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/token"
	"net/http"
	"strconv"

	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	RefreshTokens(c *gin.Context)
	GetInfo(c *gin.Context)

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

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user with the provided email and password.
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   user   body    object{username=string,email=string,password=string}  true  "User credentials"
// @Success 200 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid credentials / User already exists"
// @Router /api/auth/register [post]
func (u *userController) RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	_, accessToken, refreshToken, err := u.service.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := map[string]interface{}{
		//"user": gin.H{
		//	"id":    user.ID,
		//	"email": user.Email,
		//},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	//c.SetCookie("access_token", accessToken, 900, "/", "localhost", false, true)
	//c.SetCookie("refresh_token", refreshToken, 604800, "/", "localhost", false, true)

	c.JSON(http.StatusOK, response)
}

// Login godoc
// @Summary Login user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body object{email=string,password=string} true "User credentials"
// @Success 200 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid credentials"
// @Router /api/auth/login [post]
func (u *userController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	_, accessToken, refreshToken, err := u.service.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//c.SetCookie("access_token", accessToken, 900, "/", "localhost", false, true)
	//c.SetCookie("refresh_token", refreshToken, 604800, "/", "localhost", false, true)

	response := map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusOK, response)
}

// RefreshTokens godoc
// @Summary Get new tokens
// @Description Take refresh token and produce 2 new tokens
// @Tags Users
// @Accept json
// @Produce json
// Param token body object{refresh_token=string} true "Refresh token"
// Success 200 {object} map[string]string "Tokens generated successfully"
// @Failure 400 {object} map[string]string "Invalid token"
// @Router /api/auth/refresh [post]
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

// GetInfo godoc
// @Summary Get user info
// @Description Take access token from header  and provide user info
// @Tags Users
// @Accept json
// @Produce json
// Success 200 {object} map[string]interface{} "Info provided successfully"
// @Failure 400 {object} map[string]string "Invalid token"
// @Router /api/auth/user-info [get]
func (u *userController) GetInfo(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found in cookies"})
		return
	}

	//uid, _ := token.ParseToken(accessToken)
	//uid, _ := token.ParseToken(c.Request)

	//authHeader := c.GetHeader("Authorization")
	//
	//if authHeader == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
	//	return
	//}
	//
	//parts := strings.Split(authHeader, " ")
	//if len(parts) != 2 || parts[0] != "Bearer" {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header format"})
	//	return
	//}
	//
	//accessToken := parts[1]

	userID, err := token.VerifyToken(accessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.service.FindById(strconv.Itoa(int(userID)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"img":      user.AvatarPath,
		"phone":    user.Phone,
	})
}

// GetUsers godoc
// @Summary Get users info
// @Description Provide users info
// @Tags Users
// @Accept json
// @Produce json
// Success 200 {object} map[string]interface{} "Info provided successfully"
// @Failure 400 {object} map[string]string "Server error"
// @Router /api/users [get]
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
