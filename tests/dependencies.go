package tests

import (
	"envs/internal/core/services"
	"envs/internal/handlers"
	"envs/internal/repositories"
	"envs/pkg/database"
	"envs/pkg/validator"
)

func InitUserDependencies(db database.Connection) *handlers.User {
	return handlers.NewUser(
		services.NewUser(
			repositories.NewUser(db)),
		validator.New(),
	)
}
