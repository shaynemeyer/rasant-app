package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/shaynemeyer/rasant"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err!= nil {
    log.Fatal(err)
  }

	// init rasant
	ras := &rasant.Rasant{}
	err = ras.New(path)
	if err!= nil {
    log.Fatal(err)
  }

	ras.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: ras,
	} 

	myHandlers := &handlers.Handlers{
		App: ras,
	}

	app := &application{
		App: ras,
		Handlers: myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}