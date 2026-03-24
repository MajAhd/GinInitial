package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	// Switch to test mode to avoid unnecessary logs
	gin.SetMode(gin.TestMode)

	// Create a dummy logger that writes to a buffer for tests
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	router := setupRouter(logger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but instead got %d", http.StatusOK, w.Code)
	}

	// Check response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response["message"] != "pong" {
		t.Errorf("Expected message to be 'pong' but instead got '%s'", response["message"])
	}
}
