package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func clients(w http.ResponseWriter, r *http.Request) {
	//response, err := http.Get("http://teleaus.com/api/v1/partners")
	//
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//var result map[string]interface{}
	//json.NewDecoder(response.Body).Decode(&result)
	//

	client := http.Client{}
	request, err := http.NewRequest("GET", "http://teleaus.com/api/v1/partners", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
}

func handleRequests() {
	http.HandleFunc("/", clients)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
