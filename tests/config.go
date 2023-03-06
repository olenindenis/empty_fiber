package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"envs/pkg/database"
	"envs/pkg/logger"

	"github.com/joho/godotenv"
)

const (
	requestTimeout  = 1000
	envFileName     = ".env"
	userChangingUri = "/api/v1/user/:id"
	userListUri     = "/api/v1/user/list"
)

var (
	errEnvFileUnavailable = errors.New("env: .env file unavailable")
	DB                    database.Connection
)

func loadEnvs() {
	log := logger.New(logger.Dev)
	if _, err := os.Stat(envFileName); err == nil {
		var fileEnv map[string]string
		fileEnv, err := godotenv.Read()
		if err != nil {
			log.Warn(errEnvFileUnavailable)
		}

		for key, val := range fileEnv {
			if len(os.Getenv(key)) == 0 {
				if err := os.Setenv(key, val); err != nil {
					return
				}
			}
		}
	}
}

func loadDBConnection() {
	log := logger.New(logger.Dev)
	var err error
	DB, err = database.NewDBConnection(
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
		log.Fatal(fmt.Sprintf("connection error: %v \n", err))
	}
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	return buf, e.Encode(data)
}

func parseTime(t *testing.T, value string) time.Time {
	date, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatal(err)
	}
	return date
}

func init() {
	loadEnvs()
	loadDBConnection()
}
