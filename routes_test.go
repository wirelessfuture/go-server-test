package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducts(t *testing.T) {
	tt := []struct {
		name   string
		method string
		url    string
		status int
	}{
		{name: "Test GET method allowed", method: "GET", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusOK},
		{name: "Test PUT method allowed", method: "PUT", url: "http://192.168.0.18:8080/products?item=shoes&price=45", status: http.StatusOK},
		{name: "Test POST method allowed", method: "POST", url: "http://192.168.0.18:8080/products?item=hat&price=25", status: http.StatusOK},
		{name: "Test DELETE method allowed", method: "DELETE", url: "http://192.168.0.18:8080/products?item=hat", status: http.StatusOK},
		{name: "Test PATCH method not allowed", method: "PATCH", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusMethodNotAllowed},
		{name: "Test HEAD method not allowed", method: "HEAD", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusMethodNotAllowed},
		{name: "Test OPTIONS method not allowed", method: "OPTIONS", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusMethodNotAllowed},
		{name: "Test CONNECT method not allowed", method: "CONNECT", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusMethodNotAllowed},
		{name: "Test TRACE method not allowed", method: "TRACE", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusMethodNotAllowed},
	}

	for _, tc := range tt {
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatalf("could not create request %v", err)
		}

		rec := httptest.NewRecorder()
		db.products(rec, req)

		res := rec.Result()
		if res.StatusCode != tc.status {
			t.Errorf("%v - expected status %v; got %v", tc.name, tc.status, res.StatusCode)
		}
	}
}
