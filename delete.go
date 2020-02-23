package main

import (
	"fmt"
	"net/http"
)

// Given an item name, deletes that item
func (db dataBase) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "improper query")
		return
	}

	mu.Lock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s does not exist", item)
		return
	}
	delete(db, item)
	mu.Unlock()

	fmt.Fprintf(w, "item deleted: %s", item)
}
