package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gbmoraes-dev/blueprint/internal/respond"
)

type HealthHandler struct {
	startTime time.Time
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{startTime: time.Now()}
}

type healthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	if err := respond.WriteJSON(w, http.StatusOK, healthResponse{
		Status: "ok",
		Uptime: time.Since(h.startTime).Truncate(time.Second).String(),
	}); err != nil {
		slog.Error("failed to write response", "error", err)
		respond.WriteError(w, http.StatusInternalServerError, "internal server error")
	}
}
