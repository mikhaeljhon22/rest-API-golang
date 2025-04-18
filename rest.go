package main 
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

//get method
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Desccription", Content: "Hello World"},
	}

	json.NewEncoder(w).Encode(articles)
}



func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello Golang")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main(){
handleRequest()
}