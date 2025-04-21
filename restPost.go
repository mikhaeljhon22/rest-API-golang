package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Content struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
    Age int `json:"Age"`
}

type Contents []Content


var contents Contents = Contents{}

func newContent(w http.ResponseWriter, r *http.Request) {
	var newContent Content 
	err := json.NewDecoder(r.Body).Decode(&newContent)
	/* Decode body and request body,
	checl err if nil create http error and statusBadRequest*/
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
	}
	
	
	contents = append(contents, newContent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Data berhasil ditambahkan",
		"data":    newContent,
	}

	json.NewEncoder(w).Encode(response)
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/newContent", newContent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}