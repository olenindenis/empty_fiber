package application

import (
	"envs/pkg/cache"
	"envs/pkg/database"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

var (
	envFileUnavailable = errors.New("env: .env file unavailable")
)

const (
	envFileName = ".env"
)

func Envs() {
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

func NewLogger(levelString string) *log.Logger {
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

	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)

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

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Errorf("error env DB_PORT: %w \n", err))
	}
	dbConnection, err := database.NewDBConnection(
		database.DriverName(os.Getenv("DB_DRIVER")),
		database.DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(port),
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
