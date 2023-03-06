package requests

import (
	"envs/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

// Request example
type Request struct {
	Ctx       *fiber.Ctx
	Validator ports.Validator
}
