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

func NewApp(logLevel string) App {
	initLogs(logLevel)
	initEnv()

	dependencies := NewDependencies()

	return App{
		httpServer: NewHttpServer(
			WithHost(os.Getenv("LISTEN_HOST")),
			WithPort(os.Getenv("LISTEN_PORT")),
		),
		dependencies: dependencies,
	}
}

func initLogs(levelString string) {
	var level log.Level

	if len(levelString) == 0 {
		levelString = os.Getenv("LOG_LEVEL")
	}

	if len(levelString) == 0 {
		levelString = "error"
	}

	level, err := log.ParseLevel(levelString)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Run with log level: %s", level)

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
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
