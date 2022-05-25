package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var _ Connection = (*MySqlConnection)(nil)

type MySqlConnection struct {
	config DBConfig
}

func NewMySqlConnection(config DBConfig) Connection {
	return MySqlConnection{
		config: config,
	}
}

func (s MySqlConnection) Connection() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		s.config.Username,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.Database,
	)

	conn, err := sql.Open(string(mysqlDriver), dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s MySqlConnection) Driver() DriverName {
	return mysqlDriver
}
