package main

import (
	"fmt"
	"net/http"
)

// Read an items price given an item name, or read all items if all passed
func (db dataBase) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "improper query")
		return
	}

	if item == "all" {
		mu.Lock()
		for item, price := range db {
			fmt.Fprintf(w, "%s: %0.2f\n", item, price)
		}
		mu.Unlock()
		return
	}

	mu.Lock()
	price, ok := db[item]
	mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "$%0.2f\n", price)
}
