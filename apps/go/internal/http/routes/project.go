package project

import (
	"net/http"

	project "github.com/arunsingh28/openaxis/internal/http/handlers"
)

func RegisterProjectRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /init-project", project.InitProject)
}
