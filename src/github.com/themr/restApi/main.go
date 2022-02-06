package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Struct is just like a class
//Books Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init Book Variable as Slice Book Struct
var books []Book //This is the Slice of Book of Type Book Struct

//Get Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single bool
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //Get parameters
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Creat a new book
func createbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, items := range books {
		if items.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, items := range books {
		if items.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	fmt.Println("Server is running on port 8000")
	//Initialize Router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, Book{ID: "1", Title: "Book One", Isbn: "1223", Author: &Author{
		Firstname: "Siemen", Lastname: "Doe",
	}})
	books = append(books, Book{ID: "2", Title: "Book Two", Isbn: "2312", Author: &Author{
		Firstname: "Siemen", Lastname: "Doe",
	}})
	books = append(books, Book{ID: "3", Title: "Book Three", Isbn: "13105", Author: &Author{
		Firstname: "Siemen", Lastname: "Doe",
	}})

	//Router Handles/ Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createbook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r)) //This line is used to tell the server thath on which port you will execute

}
