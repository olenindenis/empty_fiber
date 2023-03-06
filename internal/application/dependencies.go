package application

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"envs/pkg/cache"
	"envs/pkg/database"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	envFileUnavailable = errors.New("env: .env file unavailable")
)

const (
	envFileName = ".env"
	levelForLog = "dev"
)

func Envs() {
	log := NewLogger(levelForLog).Sugar()

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

func NewLogger(level string) *zap.Logger {
	var logger *zap.Logger
	var err error
	defer logger.Sync()

	switch level {
	case "prod":
		logger, err = zap.NewProduction()
		if err != nil {
			return nil
		}
	case "dev":
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil
		}
	}

	return logger
}

func NewCacheConnection() *cache.Cache {
	cacheHost := fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"))
	mem := cache.NewCache(cacheHost)
	err := mem.Ping()
	if err != nil {
		panic(err.Error())
	}
	return mem
}

func NewDBConnection() database.Connection {
	startTimeout, err := strconv.Atoi(os.Getenv("DB_START_TIMEOUT"))
	if err != nil {
		panic(fmt.Errorf("error env DB_START_TIMEOUT: %w", err))
	}
	time.Sleep(time.Second * time.Duration(startTimeout))
	dbConnection, err := database.NewDBConnection(
		database.DriverName(os.Getenv("DB_DRIVER")),
		database.DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}
	conn, err := dbConnection.Connection()
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}
	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}

	return dbConnection
}
