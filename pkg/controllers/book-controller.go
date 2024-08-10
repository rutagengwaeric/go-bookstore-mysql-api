package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rutagengwaeric/go-bookstore-mysql-api/pkg/config"
	"github.com/rutagengwaeric/go-bookstore-mysql-api/pkg/models"
	"github.com/rutagengwaeric/go-bookstore-mysql-api/pkg/utils"
)

// GetBook handles the request to retrieve all books from the database
func GetBook(w http.ResponseWriter, r *http.Request) {
	// Retrieve all books from the database using the GetAllBooks function
	books, err := models.GetAllBooks()
	if err != nil {
		// If an error occurs, return a 500 Internal Server Error status
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}
	// Marshal the books into JSON format
	res, _ := json.Marshal(books)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the JSON response to the client
	w.Write(res)
}

// GetBookById handles the request to retrieve a specific book by its ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the request URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// Convert the book ID from string to int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If conversion fails, return a 400 Bad Request status
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Retrieve the book details by ID using the GetBookById function
	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		// If the book is not found, return a 404 Not Found status
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Marshal the book details into JSON format
	res, _ := json.Marshal(bookDetails)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the JSON response to the client
	w.Write(res)
}

// CreateBook handles the request to create a new book record
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Book struct
	var book models.Book
	if err := utils.ParseBody(r, &book); err != nil {
		// If parsing fails, return a 400 Bad Request status
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the new book record using the CreateBook function
	createdBook := book.CreateBook()
	if createdBook == nil {
		// If creation fails, return a 500 Internal Server Error status
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	// Marshal the created book into JSON format
	res, _ := json.Marshal(createdBook)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 201 Created
	w.WriteHeader(http.StatusCreated)
	// Write the JSON response to the client
	w.Write(res)
}

// DeleteBook handles the request to delete a specific book by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the request URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// Convert the book ID from string to int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If conversion fails, return a 400 Bad Request status
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Delete the book by ID using the DeleteBook function
	deletedBook, err := models.DeleteBook(ID)
	if err != nil {
		// If the book is not found or deletion fails, return a 404 Not Found status
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Marshal the deleted book into JSON format
	res, _ := json.Marshal(deletedBook)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the JSON response to the client
	w.Write(res)
}

// UpdateBook handles the request to update an existing book record
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Book struct
	var updateBook models.Book
	if err := utils.ParseBody(r, &updateBook); err != nil {
		// If parsing fails, return a 400 Bad Request status
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract the book ID from the request URL parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// Convert the book ID from string to int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		// If conversion fails, return a 400 Bad Request status
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Retrieve the book details by ID using the GetBookById function
	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		// If the book is not found, return a 404 Not Found status
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update the book fields if they are provided
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save the updated book record to the database
	if err := config.GetDB().Save(&bookDetails).Error; err != nil {
		// If saving fails, return a 500 Internal Server Error status
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	// Marshal the updated book details into JSON format
	res, _ := json.Marshal(bookDetails)
	// Set the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 200 OK
	w.WriteHeader(http.StatusOK)
	// Write the JSON response to the client
	w.Write(res)
}
