package application

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Router(server *fiber.App, dependencies Dependencies) {
	docs := server.Group("/docs")
	docs.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "sysadmin",
		},
	}))
	docs.Get("/*", swagger.HandlerDefault)

	server.Use(requestid.New())
	server.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${method} ${path} ${status} pid=${pid}\n",
		TimeFormat: DateTimeLayout,
	}))
	server.Use(recover.New())

	server.Get("/health_checks", dependencies.healthChecksHandlers.HealthChecks)

	v1 := server.Group("/api/v1")

	usersRoutes := v1.Group("/user")
	usersRoutes.Get("/list", dependencies.userHandlers.List)
	usersRoutes.Get("/:id", dependencies.userHandlers.Show)
	usersRoutes.Put("/:id", dependencies.userHandlers.Update)
	usersRoutes.Delete("/:id", dependencies.userHandlers.Delete)
}
