package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is a struct
type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles  is a arrary
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Enpoint Hit: homePage")
}

func handleRequests() {
	/*
		http.HandleFunc("/", homePage)
		http.HandleFunc("/articles", returnAllArticles)
		log.Fatal(http.ListenAndServe(":10000", nil))
	*/

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {

	Articles = []Article{
		{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{ID: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit : returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBoy, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBoy))

	var article Article
	json.Unmarshal(reqBoy, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.ID == id {
			fmt.Println(Articles[:index])
			fmt.Println(Articles[index+1:])
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
