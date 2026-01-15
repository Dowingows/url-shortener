package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	res := httptest.NewRecorder()

	pingHandler(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("esperado status 200, recebido %d", res.Code)
	}

	expected := "pong\n"
	if res.Body.String() != expected {
		t.Fatalf("esperado %q, recebido %q", expected, res.Body.String())
	}
}
