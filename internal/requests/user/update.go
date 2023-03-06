package user

import (
	"strconv"

	"envs/internal/core/ports"
	"envs/internal/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fastjson"
)

// UpdateRequest example
type UpdateRequest struct {
	requests.Request
	ID    uint   `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func NewUpdateRequest(ctx *fiber.Ctx, validator ports.Validator) *UpdateRequest {
	return &UpdateRequest{Request: requests.Request{Ctx: ctx, Validator: validator}}
}

func (sr *UpdateRequest) Validate() error {
	userID, err := strconv.Atoi(sr.Ctx.Params("id"))
	if err != nil {
		return err
	}
	sr.ID = uint(userID)

	values, err := fastjson.ParseBytes(sr.Ctx.Body())
	if err != nil {
		return err
	}

	sr.Name = string(values.GetStringBytes("name"))
	sr.Email = string(values.GetStringBytes("email"))

	return sr.Validator.Struct(sr)
}
