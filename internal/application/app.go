package application

import (
	"context"
	"os"
	"strings"
	"time"

	"envs/internal/core/ports"
	"envs/internal/core/services"
	"envs/internal/handlers"
	"envs/internal/repositories"
	"envs/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/olekukonko/tablewriter"
)

type App struct {
	httpDSN HttpDSN
	level   string
}

func NewApp(level string) App {
	Envs()

	return App{
		level: level,
		httpDSN: NewHttpDSN(
			WithHost(os.Getenv("LISTEN_HOST")),
			WithPort(os.Getenv("LISTEN_PORT")),
		),
	}
}

func (a *App) Run() {
	fx.New(
		fx.NopLogger,
		fx.Supply(a.level, a.httpDSN),
		fx.Provide(
			NewLogger,
			NewDBConnection,
			NewServer,
			NewCacheConnection,
		),
		fx.Provide(
			fx.Annotate(repositories.NewUserRepository, fx.As(new(ports.UserRepository))),
			fx.Annotate(services.NewUserService, fx.As(new(ports.UserService))),
			fx.Annotate(handlers.NewUserHandler, fx.As(new(ports.UserHandlers))),
			fx.Annotate(validator.NewValidator, fx.As(new(ports.Validator))),
			fx.Annotate(handlers.NewHealthChecksHandlers, fx.As(new(ports.HealthChecksHandlers))),
		),
		fx.Invoke(
			ConfigureMiddleware,
			DocsRotes,
			HealthChecksRoutes,
			UserRoutes,
		),
		fx.Invoke(dependencies),
	).Run()
}

func dependencies(server *fiber.App, dsn HttpDSN, lifecycle fx.Lifecycle) {
	PrintRoutes(server)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				errChan := make(chan error)

				go func() {
					errChan <- server.Listen(dsn.DSN())
				}()

				select {
				case err := <-errChan:
					return err
				case <-time.After(100 * time.Millisecond):
					return nil
				}
			},
			OnStop: func(ctx context.Context) error {
				return server.Shutdown()
			},
		},
	)
}

func PrintRoutes(server *fiber.App) {
	var arr [][]string
	for _, routes := range server.Stack() {
		for _, route := range routes {
			arr = append(arr, []string{route.Method, route.Path, route.Name, strings.Join(route.Params, ",")})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Method", "Path", "Route name", "Params"})
	table.SetBorder(false)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)

	table.SetColumnColor(
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{})

	// table.AppendBulk(arr)

	for _, row := range arr {
		if row[0] == "GET" {
			table.Rich(row, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{}})
		} else if row[0] == "POST" {
			table.Rich(row, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgGreenColor},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{}})
		} else if row[0] == "PUT" {
			table.Rich(row, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgYellowColor},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{}})
		} else if row[0] == "DELETE" {
			table.Rich(row, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.Normal, tablewriter.FgRedColor},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{}})
		} else if row[0] == "TRACE" {
			table.Rich(row, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.Normal, tablewriter.BgHiRedColor},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{}})
		} else {
			table.Append(row)
		}
	}
	table.Render()
}
