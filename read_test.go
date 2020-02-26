package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRead(t *testing.T) {
	tt := []struct {
		name   string
		method string
		url    string
		status int
	}{
		{name: "Test GET exists", method: "GET", url: "http://192.168.0.18:8080/products?item=shoes", status: http.StatusOK},
		{name: "Test GET exists", method: "GET", url: "http://192.168.0.18:8080/products?item=socks", status: http.StatusOK},
		{name: "Test GET not exists", method: "GET", url: "http://192.168.0.18:8080/products?item=shoess", status: http.StatusNotFound},
		{name: "Test GET not exists", method: "GET", url: "http://192.168.0.18:8080/products?item=shoessss", status: http.StatusNotFound},
		{name: "Test GET not exists", method: "GET", url: "http://192.168.0.18:8080/products?item=baseball", status: http.StatusNotFound},
		{name: "Test GET bad request", method: "GET", url: "http://192.168.0.18:8080/products?itemm=shoes", status: http.StatusBadRequest},
		{name: "Test GET bad request", method: "GET", url: "http://192.168.0.18:8080/products?items=shoes", status: http.StatusBadRequest},
		{name: "Test GET empty query param", method: "GET", url: "http://192.168.0.18:8080/products?item=", status: http.StatusBadRequest},
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
