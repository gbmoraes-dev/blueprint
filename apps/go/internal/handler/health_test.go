package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gbmoraes-dev/blueprint/internal/handler"
)

func TestHealth_OK(t *testing.T) {
	h := handler.NewHealthHandler()

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/health", http.NoBody)

	w := httptest.NewRecorder()

	h.Health(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	var body struct {
		Status string `json:"status"`
		Uptime string `json:"uptime"`
	}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status %q, got %q", "ok", body.Status)
	}
}

func TestHealth_InternalServerError(t *testing.T) {
	h := handler.NewHealthHandler()

	req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/health", http.NoBody)

	w := &failWriter{header: make(http.Header)}

	h.Health(w, req)

	if w.statusCode != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", w.statusCode)
	}
}

type failWriter struct {
	header     http.Header
	statusCode int
}

func (f *failWriter) Header() http.Header {
	return f.header
}

func (f *failWriter) WriteHeader(code int) {
	f.statusCode = code
}

func (f *failWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}
