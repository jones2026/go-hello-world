package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jones2026/go-hello-world/healthz"
)

func TestHealthCheckHandlerErrorThrowing(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz?errorType=FakeError", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	hc := &healthz.Config{
		Hostname: "some fake host",
	}
	handler, err := healthz.Handler(hc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestHealthCheckHandlerHappyPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	hc := &healthz.Config{
		Hostname: "some fake host",
	}
	handler, err := healthz.Handler(hc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
