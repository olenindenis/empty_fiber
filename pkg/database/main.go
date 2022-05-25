package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type DriverName string

const (
	mysqlDriver    DriverName = "mysql"
	postgresDriver            = "postgres"
)

type Connection interface {
	Connection() (*sql.DB, error)
	Driver() DriverName
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Database string
}

func NewDBConnection(driverName DriverName, config DBConfig) (Connection, error) {
	switch driverName {
	case mysqlDriver:
		return NewMySqlConnection(config), nil
	case postgresDriver:
		return NewPostgresConnection(config), nil
	}

	return nil, errors.New(fmt.Sprintf("error: driver %s not found", driverName))
}
