package middleware

import (
	"myapp/data"

	"github.com/shaynemeyer/rasant"
)

type Middleware struct {
	App *rasant.Rasant
	Models data.Models
}