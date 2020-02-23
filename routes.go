package main

import (
	"fmt"
	"net/http"
)

func (db dataBase) products(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		db.create(w, req)
	case "GET":
		db.read(w, req)
	case "PUT":
		db.update(w, req)
	case "DELETE":
		db.delete(w, req)
	default:
		fmt.Fprintf(w, "Method not allowed")
	}
}
