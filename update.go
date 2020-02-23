package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Given an item name and a valid price, updates the price of that item
func (db dataBase) update(w http.ResponseWriter, req *http.Request) {
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
	db[item] = convertedPrice
	mu.Unlock()

	fmt.Fprintf(w, "item price updated: %s to %0.2f", item, convertedPrice)
}
