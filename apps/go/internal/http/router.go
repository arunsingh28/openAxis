package router

import (
	"net/http"

	project "github.com/arunsingh28/openaxis/internal/http/routes"
)

func RegisterRoutes(mux *http.ServeMux) {
	project.RegisterProjectRoutes(mux)
}
