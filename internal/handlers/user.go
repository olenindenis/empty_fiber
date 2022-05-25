package handlers

import (
	"envs/internal/core/ports"
	userRequest "envs/internal/requests/user"
	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandlers struct {
	validator   ports.Validator
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService, validator ports.Validator) *UserHandlers {
	return &UserHandlers{
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
// @Success 200 {object} []domain.User "User domain models"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/list [get]
func (sh *UserHandlers) List(ctx *fiber.Ctx) error {
	request := userRequest.NewListRequest(ctx, sh.validator)
	err := request.Validate()

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
	}

	users, err := sh.userService.List(request.Limit, request.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
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
// @Success 200 {object} domain.User "User domain model"
// @Failure 400 {object} HTTPError "Bad request"
// @Failure 401 {object} HTTPError "Unauthorized"
// @Failure 403 {object} HTTPError "Forbidden"
// @Failure 405 {object} HTTPError "Method not allowed"
// @Failure 422 {object} HTTPError "Validation error"
// @Failure 429 {object} HTTPError "Too Many Requests"
// @Failure 500 {object} ServerError "Server error"
// @Router /api/v1/user/{id} [get]
func (sh *UserHandlers) Show(ctx *fiber.Ctx) error {
	request := userRequest.NewShowRequest(ctx, sh.validator)
	err := request.Validate()

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
	}

	user, err := sh.userService.Show(request.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
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
func (sh *UserHandlers) Update(ctx *fiber.Ctx) error {
	request := userRequest.NewUpdateRequest(ctx, sh.validator)
	err := request.Validate()

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
	}

	user, err := sh.userService.Show(request.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	user.Name = request.Name
	user.Email = request.Email
	err = sh.userService.Update(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return ctx.Status(http.StatusOK).JSON(HTTPSuccess{
		Code:    http.StatusOK,
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
func (sh *UserHandlers) Delete(ctx *fiber.Ctx) error {
	request := userRequest.NewDeleteRequest(ctx, sh.validator)
	err := request.Validate()

	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		})
	}

	err = sh.userService.Delete(request.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(HTTPError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return ctx.Status(http.StatusOK).JSON(HTTPSuccess{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}
