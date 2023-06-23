package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/shaynemeyer/rasant"
)

type application struct {
	App *rasant.Rasant 
	Handlers *handlers.Handlers
	Models data.Models
	Middleware *middleware.Middleware
}

func main() {
	r := initApplication()
	r.App.ListenAndServe()
}
