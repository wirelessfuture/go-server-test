package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducts(t *testing.T) {
	tt := struct {
		name string
		method string
		
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/products?item=shoes", serverAddress), nil)
	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()

	db.products(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}
}
