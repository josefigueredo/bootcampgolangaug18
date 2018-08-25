package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"strings"
)

var replacer = strings.NewReplacer("!", "", "?", "", "ä", "a", "á", "a", "ë", "e", "é", "e", "ï", "i", "í", "i", "ö", "o", "ó", "o", "ü", "u", "ú", "u")

func GetRoot(w http.ResponseWriter, r *http.Request) {
	// println("GetRoot")
	w.Header().Add("content-type", "application/json; charset=utf-8")
}

func GetCount(w http.ResponseWriter, r *http.Request) {
	// println("GetCount")

	var response = make(map[string]int)
	params := mux.Vars(r)
	var input, ok = params["input"]

	if ok {
		input = strings.ToLower(input)
		input = replacer.Replace(input)
		for _, r := range input {
			c := string(r)
			response[c] = response[c] + 1
		}
	}
	json.NewEncoder(w).Encode(response)
	w.Header().Add("content-type", "application/json; charset=utf-8")
}

// main function to boot up everything
func main() {
	// println("Starting...")
	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/count", GetCount).Queries("input","{input}").Methods("GET")
	log.Fatal(http.ListenAndServe(":4567", router))
}
