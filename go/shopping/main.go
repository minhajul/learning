package main

import (
	"log"
	"net/http"
)

func getResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("Hello World")

	//db := databaseConnection()
	//
	//defer db.Close()
}

func handleRequests() {
	http.HandleFunc("/", getResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
