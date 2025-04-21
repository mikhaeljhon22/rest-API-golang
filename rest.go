package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

type Data struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type Datas []Data




type Content struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
    Age int `json:"Age"`
}

type Contents []Content

type Heboh struct {
	Birth string `json:"Birth"`
	Age int `json:"Age"`
}
type Hebohs []Heboh

var articles Articles = Articles{
	{Title: "Title", Desc: "Desc", Content: "Content"},
}

func heboh(w http.ResponseWriter, r *http.Request) {
	hebohs := Hebohs{
		{Birth: "IFODHFI", Age: 25},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hebohs)
}


func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang")
}

func testPostArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test post endpoint worked")
}

func TestingMandiri(w http.ResponseWriter, r *http.Request) {
	contents := Datas{
		{Username: "username", Password: "Password", Email: "mikhaeljhon22@gmail.com"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)
}


func createArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle Article
	json.NewDecoder(r.Body).Decode(&newArticle)
	articles = append(articles,newArticle)
	json.NewEncoder(w).Encode(newArticle)	
}

	
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticle).Methods("POST")
	myRouter.HandleFunc("/testing", TestingMandiri).Methods("GET")
	myRouter.HandleFunc("/articlePost", createArticle).Methods("POST")
	myRouter.HandleFunc("/newContent", newContent).Methods("POST")
	myRouter.HandleFunc("/heboh", heboh).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
