package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rutagengwaeric/go-bookstore-mysql-api/pkg/config"
)

// Define a global variable to hold the database connection
var db *gorm.DB

// Define the Book struct which represents the book entity in the database
type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Author      string `gorm:"type:varchar(100);not null" json:"author"`
	Publication string `gorm:"type:varchar(100);not null" json:"publication"`
}

// Initialize the package by connecting to the database and setting up the schema
func init() {
	// Establish a connection to the database using the config package
	config.Connect()
	// Retrieve the database connection from the config package
	db = config.GetDB()
	// Automatically migrate the schema, creating or updating the Book table
	db.AutoMigrate(&Book{})
}

// CreateBook saves a new book record to the database and returns the created book
func (b *Book) CreateBook() *Book {
	// Attempt to create the book record in the database
	if err := db.Create(&b).Error; err != nil {
		// Log an error message if book creation fails
		log.Printf("Failed to create book: %v", err)
		// Return nil to indicate failure
		return nil
	}
	// Return the created book if successful
	return b
}

// GetAllBooks retrieves all book records from the database
func GetAllBooks() ([]Book, error) {
	// Declare a slice to hold the retrieved books
	var books []Book
	// Attempt to find all book records and store them in the slice
	if err := db.Find(&books).Error; err != nil {
		// Return an error if the retrieval fails
		return nil, err
	}
	// Return the list of books and a nil error if successful
	return books, nil
}

// GetBookById retrieves a single book record by its ID
func GetBookById(Id int64) (*Book, error) {
	// Declare a variable to hold the retrieved book
	var book Book
	// Attempt to find the book with the given ID
	result := db.Where("id = ?", Id).First(&book)
	// Check if the book was not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Return nil and a "book not found" error if the book does not exist
		return nil, errors.New("book not found")
	}
	// Return nil and any other error if there was a problem retrieving the book
	if result.Error != nil {
		return nil, result.Error
	}
	// Return the retrieved book and a nil error if successful
	return &book, nil
}

// DeleteBook removes a book record from the database by its ID
func DeleteBook(Id int64) (*Book, error) {
	// Declare a variable to hold the book to be deleted
	var book Book
	// Attempt to find the book with the given ID
	result := db.Where("id = ?", Id).First(&book)
	// Check if the book was not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Return nil and a "book not found" error if the book does not exist
		return nil, errors.New("book not found")
	}
	// Return nil and any other error if there was a problem retrieving the book
	if result.Error != nil {
		return nil, result.Error
	}

	// Attempt to delete the book record from the database
	if err := db.Delete(&book).Error; err != nil {
		// Return nil and the error if deletion fails
		return nil, err
	}
	// Return the deleted book and a nil error if successful
	return &book, nil
}
