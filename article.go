package main

// Article represents an article with title, description and content
type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles will hold the collection
var Articles []Article
