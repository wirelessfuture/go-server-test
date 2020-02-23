package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "192.168.0.18/products?items=shoes", nil)
	if err != nil {
		t.Fatalf("could not create requestL %v", err)
	}

	rec := httptest.NewRecorder()

	db.products(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}
}
