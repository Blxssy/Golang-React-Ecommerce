package controller

import (
	"fmt"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/migration"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/test"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router, ctr = test.PrepareForControllerTest()

func TestUserController_GetUserByID(t *testing.T) {
	//router, ctr := test.PrepareForControllerTest()

	user := NewUserController(ctr)
	router.GET(config.APIUsersID, user.GetUserByID)

	migration.InitData(ctr)

	uri := request.NewRequestBuilder().URL(config.APIUsers).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{
    "email": "test@test.com",
    "img": "https://api.multiavatar.com/",
    "username": "test"
	}`,
		rec.Body.String())
}

func TestUserController_RegisterUser(t *testing.T) {
	//router, ctr := test.PrepareForControllerTest()

	user := NewUserController(ctr)
	router.POST(config.APIREGISTER, user.RegisterUser)

	param := createUserForRegister()
	req := test.NewJSONRequest("POST", config.APIREGISTER, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestUserController_LoginUser(t *testing.T) {
	//router, ctr := test.PrepareForControllerTest()

	user := NewUserController(ctr)
	router.POST(config.APILOGIN, user.Login)

	param := createUserForLogin()
	req := test.NewJSONRequest("POST", config.APILOGIN, param)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestUserController_GetUsers(t *testing.T) {
	//router, ctr := test.PrepareForControllerTest()

	user := NewUserController(ctr)
	router.GET(config.APIUsers, user.GetUsers)

	req := httptest.NewRequest("GET", config.APIUsers, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

//func TestUserController_RefreshToken(t *testing.T) {
//	router, ctr := test.PrepareForControllerTest()
//
//	user := NewUserController(ctr)
//	router.POST(config.APIREFRESH, user.RefreshTokens)
//
//	param := createRefreshToken()
//
//	req := test.NewJSONRequest("POST", config.APIREFRESH, param)
//	rec := httptest.NewRecorder()
//
//	router.ServeHTTP(rec, req)
//
//	assert.Equal(t, http.StatusOK, rec.Code)
//	assert.NotEmpty(t, rec.Body.String())
//}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type RegisterUserForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createRefreshToken() *RefreshToken {
	return &RefreshToken{
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MjUzNTM3MDJ9.RKct4O4JJegsrSWDW6f9r6VLP3pdJ10POuXpQeRv4JA",
	}
}

func createUserForRegister() *RegisterUserForm {
	return &RegisterUserForm{
		Email:    "test3@test.com",
		Password: "test",
	}
}

func createUserForLogin() *RegisterUserForm {
	return &RegisterUserForm{
		Email:    "test@test.com",
		Password: "test",
	}
}
