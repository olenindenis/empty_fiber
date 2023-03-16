package ports

import (
	"envs/internal/core/domain"
	"envs/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers interface {
	List(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type UserService interface {
	List(filter dto.ListFilter) ([]domain.User, error)
	Show(id uint) (domain.User, error)
	Update(user domain.User) error
	Delete(id uint) error
}

type UserRepository interface {
	Store(name, email, password string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Find(id uint) (domain.User, error)
	List(filter dto.ListFilter) ([]domain.User, error)
	Update(user domain.User) error
	Delete(id uint) error
}
