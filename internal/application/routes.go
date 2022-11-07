package application

import (
	"envs/internal/core/ports"
	"envs/internal/handlers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"net/http"
)

const (
	DateTimeLayout = "15:04:05 02-01-2006"
)

func ConfigureMiddleware(server *fiber.App) {
	server.Use(requestid.New())
	server.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${method} ${path} ${status} pid=${pid}\n",
		TimeFormat: DateTimeLayout,
	}))
	server.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	server.All("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNoContent).JSON(handlers.HTTPError{
			Message: http.StatusText(http.StatusNoContent),
		})
	}).Name("Root")
}

func DocsRotes(server *fiber.App) {
	docs := server.Group("/docs")
	docs.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "sysadmin",
		},
	}))
	docs.Get("/*", swagger.HandlerDefault).Name("Docs")
}

func HealthChecksRoutes(server *fiber.App, handlers ports.HealthChecksHandlers) {
	server.Get("/health_checks", handlers.HealthChecks).Name("HealthChecks")
}

func UserRoutes(server *fiber.App, handlers ports.UserHandlers) {
	v1 := server.Group("/api/v1")

	usersRoutes := v1.Group("/user")
	usersRoutes.Get("/list", handlers.List).Name("UserList")
	usersRoutes.Get("/:id", handlers.Show).Name("UserShow")
	usersRoutes.Put("/:id", handlers.Update).Name("UserUpdate")
	usersRoutes.Delete("/:id", handlers.Delete).Name("UserDelete")
}
