package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHealthHandler(t *testing.T) {
	router := chi.NewRouter()

	router.Get("/health", healthHandler)
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := "application/json"
	if contentType != w.Header().Get("Content-Type") {
		t.Errorf("Expected content type %s, got %s", contentType, w.Header().Get("Content-Type"))
	}

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}
	if response["status"] != "healthy" {
		t.Errorf("Expected status 'healthy', got %v", response["status"])
	}
}