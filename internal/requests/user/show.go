package user

import (
	"envs/internal/core/ports"
	"envs/internal/requests"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// ShowRequest example
type ShowRequest struct {
	requests.Request
	ID uint `json:"id" validate:"required"`
}

func NewShowRequest(ctx *fiber.Ctx, validator ports.Validator) *ShowRequest {
	return &ShowRequest{Request: requests.Request{Ctx: ctx, Validator: validator}}
}

func (sr *ShowRequest) Validate() error {
	userID, err := strconv.Atoi(sr.Ctx.Params("id"))
	if err != nil {
		return err
	}
	sr.ID = uint(userID)

	return sr.Validator.Struct(sr)
}
