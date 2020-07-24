package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	Name     string
	Quantity int
	Size     []string
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product := Product{
		Name:     "T-Shirt",
		Quantity: 20,
		Size:     []string{"S", "M", "L", "XL"},
	}

	json.NewEncoder(w).Encode(product)
}

func handleRequest() {
	http.HandleFunc("/", getProduct)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
