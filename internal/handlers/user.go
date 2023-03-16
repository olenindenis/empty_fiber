package handlers

import (
	"net/http"
	"strconv"

	"envs/internal/core/ports"
	"envs/internal/dto"
	userRequest "envs/internal/requests/user"

	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

const (
	defaultLimit = 50
)

type UserHandler struct {
	validator   ports.Validator
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandler)(nil)

func NewUserHandler(userService ports.UserService, validator ports.Validator) *UserHandler {
	return &UserHandler{
		validator:   validator,
		userService: userService,
	}
}

// List godoc
// @Tags user
// @Summary List users
// @Description List users
// @Accept  json
// @Produce  json
// @Param limit body uint false "Limit"
// @Param offset body uint false "Offset"
// @Success 200 {object} []domain.User "UserHandler domain models"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/list [get]
func (sh *UserHandler) List(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	order := ctx.Query("order")
	sortBy := ctx.Query("sortBy")
	listFilter := dto.NewListFilter(uint(limit), uint(offset), dto.Order(order), sortBy)

	users, err := sh.userService.List(listFilter)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(users)
}

// Show godoc
// @Tags user
// @Summary Show user
// @Description Show user
// @Accept  json
// @Produce  json
// @Param id path uint true "userID"
// @Success 200 {object} domain.User "UserHandler domain model"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/{id} [get]
func (sh *UserHandler) Show(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	user, err := sh.userService.Show(uint(userID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
	})
}

// Update godoc
// @Tags user
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param id path uint true "userID"
// @Success 200 {object} HTTPSuccess "ok"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/{id} [put]
func (sh *UserHandler) Update(ctx *fiber.Ctx) error {
	request := userRequest.NewUpdateRequest(ctx, sh.validator)
	err := request.Validate()

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	user, err := sh.userService.Show(request.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	user.Name = request.Name
	user.Email = request.Email
	err = sh.userService.Update(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(HTTPSuccess{
		Message: http.StatusText(http.StatusOK),
	})
}

// Delete godoc
// @Tags user
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param id path uint true "userID"
// @Success 200 {object} HTTPSuccess "ok"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/{id} [delete]
func (sh *UserHandler) Delete(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	err = sh.userService.Delete(uint(userID))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(HTTPSuccess{
		Message: http.StatusText(http.StatusOK),
	})
}
