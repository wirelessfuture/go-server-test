package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Given a valid item and price, creates the new item
func (db dataBase) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" || price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "improper query")
		return
	}

	convertedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid price")
		return
	}

	mu.Lock()
	_, ok := db[item]
	if !ok {
		db[item] = convertedPrice
		mu.Unlock()
		fmt.Fprintf(w, "new item added: %s - $%0.2f", item, convertedPrice)
		return
	}
	mu.Unlock()
	fmt.Fprintf(w, "item %s: %0.2f already exists!", item, convertedPrice)
	return
}
