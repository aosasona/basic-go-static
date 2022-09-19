package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	response["message"] = "Hello, World!"
	response["status"] = "success"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/json", HandleJSON)

	port := ":8080"

	log.Printf("Listening on port %v...", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
