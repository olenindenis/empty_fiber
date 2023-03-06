package database

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"envs/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"go.uber.org/zap"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationFilesPath = "database/migrations"

type Migrator struct {
	migration *migrate.Migrate
	log       *zap.SugaredLogger
}

func NewMigrator() (Migrator, error) {
	startTimeout, err := strconv.Atoi(os.Getenv("DB_START_TIMEOUT"))
	if err != nil {
		panic(fmt.Sprintf("no timeout error: %v \n", err))
	}
	time.Sleep(time.Second * time.Duration(startTimeout))
	db, err := NewDBConnection(
		DriverName(os.Getenv("DB_DRIVER")),
		DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_DATABASE"),
		},
	)
	if err != nil {
		return Migrator{}, fmt.Errorf("db connection error: %v \n", err)
	}

	driver, err := getDriver(db)
	if err != nil {
		return Migrator{}, fmt.Errorf("instance error: %v \n", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationFilesPath),
		string(db.Driver()),
		driver,
	)
	if err != nil {
		return Migrator{}, fmt.Errorf("instance error: %v \n", err)
	}

	return Migrator{
		migration: migrator,
		log:       logger.New(logger.Dev),
	}, nil
}

func getDriver(db Connection) (database.Driver, error) {
	conn, err := db.Connection()
	if err != nil {
		panic(fmt.Sprintf("connection error: %v \n", err))
	}

	switch db.Driver() {
	case mysqlDriver:
		return mysql.WithInstance(conn, &mysql.Config{})
	case postgresDriver:
		return postgres.WithInstance(conn, &postgres.Config{})
	}
	return nil, errors.New(fmt.Sprintf("error: driver %s unknown", db.Driver()))
}

func (m *Migrator) Up() error {
	if err := m.migration.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate up error: %v \n", err)
		}
	}

	m.log.Info("Migration success")

	return nil
}

func (m *Migrator) Down() error {
	if err := m.migration.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate up error: %v \n", err)
		}
	}

	m.log.Info("Migration success")

	return nil
}

func (m *Migrator) New() error {
	reader := bufio.NewReader(os.Stdin)
	println("Enter new migration name:")

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	timePoint := time.Now().Unix()
	upFileName := fmt.Sprintf("%s/%d_create_%s_table.up.sql", migrationFilesPath, timePoint, text)
	downFileName := fmt.Sprintf("%s/%d_drop_%s_table.down.sql", migrationFilesPath, timePoint, text)

	emptyFile, err := os.Create(upFileName)
	if err != nil {
		return err
	}

	err = emptyFile.Close()
	if err != nil {
		return err
	}

	emptyFile, err = os.Create(downFileName)
	if err != nil {
		return err
	}

	err = emptyFile.Close()
	if err != nil {
		return err
	}

	m.log.Info(fmt.Sprintf("Migrations (%s, %s) created", upFileName, downFileName))

	return nil
}

func (m *Migrator) GetMigrator() *migrate.Migrate {
	return m.migration
}
