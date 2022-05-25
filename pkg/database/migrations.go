package database

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

const migrationFilesPath = "database/migrations"

type Migrator struct {
	migration *migrate.Migrate
}

func NewMigrator() (Migrator, error) {
	startTimeout, err := strconv.Atoi(os.Getenv("DB_START_TIMEOUT"))
	if err != nil {
		panic(fmt.Sprintf("no timeout error: %v \n", err))
	}
	time.Sleep(time.Second * time.Duration(startTimeout))

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(fmt.Sprintf("error port number: %v \n", err))
	}
	db, err := NewDBConnection(
		DriverName(os.Getenv("DB_DRIVER")),
		DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(port),
			Database: os.Getenv("DB_DATABASE"),
		},
	)

	//if err != nil {
	//	return Migrator{}, fmt.Errorf("instance error: %v \n", err)
	//}

	driver, err := getDriver(db)
	//driver, err := mysql.WithInstance(conn, &mysql.Config{})
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

	return Migrator{migration: migrator}, nil
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

	log.Info("Migration success")

	return nil
}

func (m *Migrator) Down() error {
	if err := m.migration.Down(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate up error: %v \n", err)
		}
	}

	log.Info("Migration success")

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

	log.Println(fmt.Sprintf("Migrations (%s, %s) created", upFileName, downFileName))

	return nil
}

func (m *Migrator) GetMigrator() *migrate.Migrate {
	return m.migration
}
