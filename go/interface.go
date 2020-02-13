package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	//response, err := http.Get("https://httpbin.org/get")
	//
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//
	//body, err := ioutil.ReadAll(response.Body)
	//
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//
	//log.Println(string(body))

	formData := url.Values{
		"name": {"Minhajul"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}

