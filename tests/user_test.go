package tests

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"envs/pkg/response"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	handler := InitUserDependencies(DB)

	app := fiber.New()
	app.Get(userChangingUri, handler.Show)

	var id = 1
	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(id), 1)
	req := httptest.NewRequest(http.MethodGet, route, nil)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestList(t *testing.T) {
	handler := InitUserDependencies(DB)

	app := fiber.New()
	app.Get(userListUri, handler.List)

	payload := map[string]interface{}{
		"limit":  0,
		"offset": 0,
	}

	buf, err := encodeData(payload)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, userListUri, buf)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestUpdate(t *testing.T) {
	handler := InitUserDependencies(DB)

	app := fiber.New()
	app.Put(userChangingUri, handler.Update)

	payload := map[string]interface{}{
		"name":  faker.Name(),
		"email": faker.Email(),
	}

	buf, err := encodeData(payload)
	if err != nil {
		t.Fatal(err)
	}

	var id = 1
	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(id), 1)
	req := httptest.NewRequest(http.MethodPut, route, buf)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}

func TestUpdateValidation(t *testing.T) {
	handler := InitUserDependencies(DB)

	app := fiber.New()
	app.Put(userChangingUri, handler.Update)

	payload := map[string]interface{}{
		"name": faker.Name(),
	}

	buf, err := encodeData(payload)
	if err != nil {
		t.Fatal(err)
	}

	var id = 1
	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(id), 1)
	req := httptest.NewRequest(http.MethodPut, route, buf)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, response.Body(resp))
}

func TestDelete(t *testing.T) {
	id := 1
	handler := InitUserDependencies(DB)

	app := fiber.New()
	app.Delete(userChangingUri, handler.Delete)

	route := strings.Replace(userChangingUri, ":id", strconv.Itoa(id), 1)
	req := httptest.NewRequest(http.MethodDelete, route, nil)
	resp, _ := app.Test(req, requestTimeout)

	assert.Equal(t, http.StatusOK, resp.StatusCode, response.Body(resp))
}
