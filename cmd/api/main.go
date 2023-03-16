package main

import (
	"flag"

	"envs/internal/application"
	"envs/pkg/formater"
)

// @title Template Fiber API
// @version 0.1
// @description This is an API for Template Fiber service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	help := flag.Bool("help", false, "display help")
	h := flag.Bool("h", false, "display help")
	var level string
	flag.StringVar(&level, "level", "info", "Run server with some logging level '--level info'")
	flag.Parse()

	if *help || *h {
		formater.Colorize(formater.ColorGreen, "-h - help")
		formater.Colorize(formater.ColorGreen, "-level - Run server with some logging level '--level info'\n"+
			"Levels: trace,debug,info,warning,error,fatal,panic")
		return
	}

	app := application.NewApp(level)
	app.Run()
}
