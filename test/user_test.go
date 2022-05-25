package main

import (
	"bytes"
	"encoding/json"
	"envs/internal/core/services"
	"envs/internal/handlers"
	"envs/internal/mocks"
	"envs/pkg/response"
	"envs/pkg/validator"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUserShowByID(t *testing.T) {
	userID := uint(1)

	userService := services.NewUserService(mocks.NewUserRepository(userID))
	userHandler := handlers.NewUserHandlers(userService, validator.NewValidator())

	app := fiber.New()
	app.Get(userChangingUri, userHandler.Show)

	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(int(userID)), 1)
	req := httptest.NewRequest(http.MethodGet, route, nil)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestUserList(t *testing.T) {
	userID := uint(1)

	userService := services.NewUserService(mocks.NewUserRepository(userID))
	userHandler := handlers.NewUserHandlers(userService, validator.NewValidator())

	app := fiber.New()
	app.Get(userListUri, userHandler.List)

	payload := map[string]interface{}{
		"limit":  10,
		"offset": 0,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodGet, userListUri, bytes.NewReader(body))
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestUserUpdate(t *testing.T) {
	userID := uint(1)

	userService := services.NewUserService(mocks.NewUserRepository(userID))
	userHandler := handlers.NewUserHandlers(userService, validator.NewValidator())

	app := fiber.New()
	app.Put(userChangingUri, userHandler.Update)

	payload := map[string]interface{}{
		"name":  faker.Name(),
		"email": faker.Email(),
	}
	body, _ := json.Marshal(payload)

	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(int(userID)), 1)
	req := httptest.NewRequest(http.MethodPut, route, bytes.NewReader(body))
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestUserUpdateValidation(t *testing.T) {
	userID := uint(1)

	userService := services.NewUserService(mocks.NewUserRepository(userID))
	userHandler := handlers.NewUserHandlers(userService, validator.NewValidator())

	app := fiber.New()
	app.Put(userChangingUri, userHandler.Update)

	payload := map[string]interface{}{
		"name": faker.Name(),
	}
	body, _ := json.Marshal(payload)

	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(int(userID)), 1)
	req := httptest.NewRequest(http.MethodPut, route, bytes.NewReader(body))
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, response.Body(resp))
}

func TestUserDelete(t *testing.T) {
	userID := uint(1)

	userService := services.NewUserService(mocks.NewUserRepository(userID))
	userHandler := handlers.NewUserHandlers(userService, validator.NewValidator())

	app := fiber.New()
	app.Delete(userChangingUri, userHandler.Delete)

	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(int(userID)), 1)
	req := httptest.NewRequest(http.MethodDelete, route, nil)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}
