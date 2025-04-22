package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Contents struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email string `json:"Email"`
}


func username(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")
	content := Contents{
			Username: "scorpion",
			Password: "testing123",
			Email: "Manoko@gmail.com",
		}
	json.NewEncoder(w).Encode(content)
}

func post(w http.ResponseWriter, r *http.Request){
	var newContent Contents
	err := json.NewDecoder(r.Body).Decode(&newContent)
	if err != nil{
		http.Error(w, "Invalid request error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Data berhasil ditambahkan",
		"success": "OK",
		"data": newContent,
	}

	json.NewEncoder(w).Encode(response)
}
func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/testing", username).Methods("GET")
	myRouter.HandleFunc("/post", post).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	handleRequest()
}