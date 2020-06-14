package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PORT where the server will listen
const PORT = ":8080"

func server() {
	// Creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	// Router Controllers
	router.HandleFunc("/", homeHandler)
	// Articles Controllers
	router.HandleFunc("/articles", returnAllArticles)
	// Article Controllers
	router.HandleFunc("/article/", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", returnSingleArticle)
	// Server Setup
	fmt.Printf("Server will listen @ %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homeHandler")
	fmt.Fprintf(w, "Welcome to Home")
}

// Articles Handler
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: return all articles")
	json.NewEncoder(w).Encode(Articles)
}

// Article Handlers
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:i], Articles[i+1])
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, _ := ioutil.ReadAll(r.Body)

	var updatedArticle Article
	json.Unmarshal(body, &updatedArticle)

	for i, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:i], updatedArticle, Articles[i+1])
		}
	}
}

// Entry-point
func main() {
	Articles = []Article{
		Article{
			ID:      "1",
			Title:   "Hello",
			Desc:    "Article Description",
			Content: "Article Content"},
		Article{
			ID:      "2",
			Title:   "Hello 2",
			Desc:    "Article Description",
			Content: "Article Content"},
	}
	server()
}
