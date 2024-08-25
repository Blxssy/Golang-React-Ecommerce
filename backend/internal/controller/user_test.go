package controller

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/config"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/test"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/utils/request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_GetUserByID(t *testing.T) {
	router, container := test.PrepareForControllerTest()

	user := NewUserController(container)
	router.POST(config.APIREGISTER, user.RegisterUser)

	setUpTestData(container)

	uri := request.NewRequestBuilder().URL(config.APIUsers).PathParams("1").Build().GetRequestURL()
	req := httptest.NewRequest("GET", uri, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"data":{"id":1,"username":"test"}}`, rec.Body.String())
}

func setUpTestData(container container.Container) {
	entity := models.NewUserWithPlainPassword("test", "test@test.com", "test")
	repo := container.GetRepository()
	_, _ = entity.Create(repo)

	entity = models.NewUserWithPlainPassword("test2", "test2@test.com", "test")
	_, _ = entity.Create(repo)
}
