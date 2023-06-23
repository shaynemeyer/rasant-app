package handlers

import (
	"myapp/data"
	"net/http"
	"time"

	"github.com/shaynemeyer/rasant"
)

type Handlers struct {
	App *rasant.Rasant
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) { 
	defer h.App.LoadTime(time.Now())
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err!= nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
