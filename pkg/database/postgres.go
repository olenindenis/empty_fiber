package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	"os"
)

var _ Connection = (*PostgresConnection)(nil)

type PostgresConnection struct {
	config DBConfig
	pool   *pgxpool.Pool
	conn   *sql.DB
}

func NewPostgresConnection(config DBConfig) Connection {
	pgxConfig, err := pgx.ParseConfig(os.Getenv("PGX_DATABASE"))
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}
	connection := PostgresConnection{
		conn: stdlib.OpenDB(*pgxConfig),
	}
	connection.config = config
	//_, err := connection.ConnectionPool()
	//if err != nil {
	//	panic(fmt.Sprintf("connection pool error: %v \n", err))
	//}
	return connection
}

func (s PostgresConnection) ConnectionPool() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		s.config.Username,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.Database,
	)

	log.Info(dsn)

	var err error
	s.pool, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	//db, err := sql.Open("pgx", dsn)
	//if err != nil {
	//	return nil, err
	//}

	//db.SetMaxOpenConns()

	return s.pool, nil
}

func (s PostgresConnection) ConnectionFromPool() (*pgxpool.Conn, error) {
	pool, err := s.ConnectionPool()
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (s PostgresConnection) Connection() (*sql.DB, error) {
	return s.conn, nil
}

func (s PostgresConnection) Driver() DriverName {
	return postgresDriver
}
