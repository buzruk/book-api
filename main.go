package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var books = []Book{}

// var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Book 1", Author: "Author 1", Year: 2000},
		Book{ID: 2, Title: "Book 2", Author: "Author 2", Year: 2001},
		Book{ID: 3, Title: "Book 3", Author: "Author 3", Year: 2002},
		Book{ID: 4, Title: "Book 4", Author: "Author 4", Year: 2003},
		Book{ID: 5, Title: "Book 5", Author: "Author 5", Year: 2004},
		Book{ID: 6, Title: "Book 6", Author: "Author 6", Year: 2005},
		Book{ID: 7, Title: "Book 7", Author: "Author 7", Year: 2006},
		Book{ID: 8, Title: "Book 8", Author: "Author 8", Year: 2007},
		Book{ID: 9, Title: "Book 9", Author: "Author 9", Year: 2008})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	//router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	//http.ListenAndServe(":8000", router)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books) // returns JSON
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			return
		}
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return
	}
	log.Println(book)
	book.ID = books[len(books)-1].ID + 1
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return
	}
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
			item = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			return
		}
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
