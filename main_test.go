package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	res := httptest.NewRecorder()

	helloHandler(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("esperado status 200, recebido %d", res.Code)
	}

	expected := "Hello, World!\n"
	if res.Body.String() != expected {
		t.Fatalf("esperado %q, recebido %q", expected, res.Body.String())
	}
}
