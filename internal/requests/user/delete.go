package user

import (
	"envs/internal/core/ports"
	"envs/internal/requests"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// DeleteRequest example
type DeleteRequest struct {
	requests.Request
	ID uint `json:"id" validate:"required"`
}

func NewDeleteRequest(ctx *fiber.Ctx, validator ports.Validator) *DeleteRequest {
	return &DeleteRequest{Request: requests.Request{Ctx: ctx, Validator: validator}}
}

func (sr *DeleteRequest) Validate() error {
	userID, err := strconv.Atoi(sr.Ctx.Params("id"))
	if err != nil {
		return err
	}
	sr.ID = uint(userID)

	return sr.Validator.Struct(sr)
}
