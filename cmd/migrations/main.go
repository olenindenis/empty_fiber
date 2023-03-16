package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"envs/pkg/database"
	"envs/pkg/logger"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

var (
	envFileUnavailable = errors.New("env: .env file unavailable")
)

const (
	envFileName = ".env"
)

func initEnv(log *zap.SugaredLogger) {
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

func main() {
	log, err := logger.New(logger.Dev, os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(fmt.Sprintf("create logger error: %v \n", err))
	}

	help := flag.Bool("help", false, "display help")
	h := flag.Bool("h", false, "display help")
	migrateNew := flag.Bool("migrate:new", false, "create new migration file")
	migrateUp := flag.Bool("migrate:up", false, "run up migration")
	migrateDown := flag.Bool("migrate:down", false, "run down migration")
	migrateVersion := flag.Bool("migrate:version", false, "show current migration version")
	var version int
	flag.IntVar(&version, "migrate:force", 0, "force set migration version '--migrate:force 4'")
	flag.Parse()

	if *help || *h {
		colorize(ColorGreen, "-migrate:new - create new migration file")
		colorize(ColorGreen, "-migrate:up - run up migration")
		colorize(ColorGreen, "-migrate:down - run down migration")
		colorize(ColorGreen, "-migrate:force 4 - set migration version to n number")
		colorize(ColorGreen, "-migrate:version - show current migration version")
		return
	}

	if len(os.Args) > 1 {
		initEnv(log)
		migrator, err := database.NewMigrator(log)
		if err != nil {
			log.Fatal(err.Error())
		}
		m := migrator.GetMigrator()

		if *migrateNew {
			err = migrator.New()
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}

		if *migrateUp {
			err = migrator.Up()
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}

		if *migrateDown {
			err = migrator.Down()
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}

		if *migrateVersion {
			v, d, err := m.Version()
			if err != nil {
				log.Fatal(err.Error())
			}
			println(fmt.Sprintf("Version: %d, dirty: %v", v, d))
			return
		}

		if version >= 0 {
			err = m.Force(version)
			if err != nil {
				log.Fatal(err.Error())
			}
			return
		}
	}
}

func colorize(color Color, message string) {
	fmt.Println(string(color), message, ColorReset)
}
