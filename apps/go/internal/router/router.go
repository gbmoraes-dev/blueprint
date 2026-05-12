package router

import (
	"net/http"

	"github.com/gbmoraes-dev/blueprint/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	health := handler.NewHealthHandler()

	mux.HandleFunc("GET /health", health.Health)

	return mux
}
