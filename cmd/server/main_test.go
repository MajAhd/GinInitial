package main

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"gininitial/internal/api"

	"github.com/gin-gonic/gin"
)

func TestPingRouteV1(t *testing.T) {
	// Prevent unnecessary Gin mode logging noise
	gin.SetMode(gin.TestMode)

	// Suppress standard log output by pushing it to a buffer
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	router := api.SetupRouter(logger)

	w := httptest.NewRecorder()
	// Note: We use the v1 path since versions are grouped in api/router.go
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but instead got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	if response["message"] != "pong" {
		t.Errorf("Expected message 'pong' but got '%s'", response["message"])
	}
	if response["version"] != "v1" {
		t.Errorf("Expected version 'v1' but got '%s'", response["version"])
	}
}
