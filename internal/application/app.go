package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	envFileUnavailable = errors.New("env: .env file unavailable")
)

const (
	envFileName = ".env"
)

type App struct {
	httpServer   *HttpServer
	dependencies Dependencies
}

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	initEnv()
}

// NewApp @title Envs api
// @version 1.1
// @description This is an API for Envs service
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func NewApp() App {
	dependencies := NewDependencies()

	return App{
		httpServer: NewHttpServer(
			WithHost(os.Getenv("LISTEN_HOST")),
			WithPort(os.Getenv("LISTEN_PORT")),
		),
		dependencies: dependencies,
	}
}

func initEnv() {
	if _, err := os.Stat(envFileName); err == nil {
		var fileEnv map[string]string
		fileEnv, err := godotenv.Read()
		if err != nil {
			log.Warn(envFileUnavailable)
		}

		for key, val := range fileEnv {
			if len(os.Getenv(key)) == 0 {
				os.Setenv(key, val)
			}
		}
	}
}

func (a *App) RunServer() {
	httpServer := a.httpServer.GetServer()
	Router(httpServer, a.dependencies)

	go func() {
		if err := httpServer.Listen(a.httpServer.GetDSN()); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	_ = httpServer.Shutdown()
}

func (a *App) CleanupTasks() {
	conn, err := a.dependencies.dbConnection.Connection()
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Close application")
}

func (a *App) RouteList() {
	app := a.httpServer.GetServer()
	Router(app, a.dependencies)
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))
}
