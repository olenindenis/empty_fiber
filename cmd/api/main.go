package main

import (
	"envs/internal/application"
)

func main() {
	app := application.NewApp()
	app.RunServer()
	// Your cleanup tasks go here
	app.CleanupTasks()
}
