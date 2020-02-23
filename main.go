package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type dataBase map[string]float64

var (
	serverAddress = "192.168.0.18:3000"
	mu            sync.Mutex
	db            = dataBase{"shoes": 50, "socks": 5}
)

func main() {

	http.HandleFunc("/products", db.products)

	fmt.Printf("Server now listening on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
