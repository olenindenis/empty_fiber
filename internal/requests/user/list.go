package user

import (
	"envs/internal/core/ports"
	"envs/internal/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fastjson"
)

const (
	defaultLimit = 50
)

// ListRequest example
type ListRequest struct {
	requests.Request
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}

func NewListRequest(ctx *fiber.Ctx, validator ports.Validator) *ListRequest {
	return &ListRequest{Request: requests.Request{Ctx: ctx, Validator: validator}}
}

func (sr *ListRequest) Validate() error {
	values, err := fastjson.ParseBytes(sr.Ctx.Body())
	if err != nil {
		return err
	}

	sr.Limit = values.GetUint("limit")
	if sr.Limit == 0 {
		sr.Limit = defaultLimit
	}
	sr.Offset = values.GetUint("offset")

	return sr.Validator.Struct(sr)
}
