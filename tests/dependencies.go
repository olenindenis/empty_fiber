package tests

import (
	"envs/internal/core/services"
	"envs/internal/handlers"
	"envs/internal/repositories"
	"envs/pkg/database"
	"envs/pkg/validator"
)

func InitUserDependencies(db database.Connection) *handlers.UserHandler {
	return handlers.NewUserHandler(
		services.NewUserService(
			repositories.NewUserRepository(db)),
		validator.NewValidator(),
	)
}
