package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	// ignore error _ because this is a well formed request
	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "status code should be 200")
	assert.Equal(t, "pong", w.Body.String(), "incorrect return value")
}
