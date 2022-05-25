package application

import (
	"envs/internal/core/ports"
	"envs/internal/core/services"
	"envs/internal/handlers"
	"envs/internal/repositories"
	"envs/pkg/cache"
	"envs/pkg/database"
	"envs/pkg/validator"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Dependencies struct {
	cache                *cache.Cache
	dbConnection         database.Connection
	healthChecksHandlers ports.HealthChecksHandlers
	userHandlers         ports.UserHandlers
}

func NewDependencies() Dependencies {
	cacheInstance := newCacheConnection()
	dbConnection := newDBConnection()
	reqValidator := newValidator()

	userHandlers := newUserHandlers(dbConnection, reqValidator)

	return Dependencies{
		cache:                cacheInstance,
		dbConnection:         dbConnection,
		healthChecksHandlers: newHealthChecksHandlers(cacheInstance, dbConnection),
		userHandlers:         userHandlers,
	}
}

func newHealthChecksHandlers(cache *cache.Cache, dbConnection database.Connection) ports.HealthChecksHandlers {
	return handlers.NewHealthChecksHandlers(cache, dbConnection)
}

func newCacheConnection() *cache.Cache {
	cacheHost := fmt.Sprintf("%s:%s", os.Getenv("CACHE_HOST"), os.Getenv("CACHE_PORT"))
	mem := cache.NewCache(cacheHost)
	err := mem.Ping()
	if err != nil {
		panic(err.Error())
	}
	return mem
}

func newDBConnection() database.Connection {
	startTimeout, err := strconv.Atoi(os.Getenv("DB_START_TIMEOUT"))
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}
	time.Sleep(time.Second * time.Duration(startTimeout))

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Sprintf("error port number: %v \n", err))
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

func newValidator() ports.Validator {
	return validator.NewValidator()
}

func newUserHandlers(dbConnection database.Connection, validator ports.Validator) ports.UserHandlers {
	userRepository := repositories.NewUserRepository(dbConnection)
	usersService := services.NewUserService(userRepository)
	return handlers.NewUserHandlers(usersService, validator)
}
