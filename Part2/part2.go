package main

import (
	"fmt"
)

// Book struct with basic details
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// Method to initialize a new book
func NewBook(title, author, isbn string, available bool) Book {
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: available,
	}
}

// Method to display book details
func (b Book) DisplayDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, Available: %v", b.Title, b.Author, b.ISBN, b.Available)
}

// EBook struct which embeds Book struct and adds a file size
type EBook struct {
	Book
	FileSize int // File size in MB
}

// Override the DisplayDetails method for EBook
func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, Available: %v, FileSize: %dMB", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}

// BookInterface interface with DisplayDetails method
type BookInterface interface {
	DisplayDetails() string
}

// Library struct to manage collection of BookInterface (Book and EBook)
type Library struct {
	Books []BookInterface
}

// Add a book or ebook to the library
func (l *Library) AddBook(book BookInterface) {
	l.Books = append(l.Books, book)
}

// List all books and ebooks in the library
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available.")
		return
	}
	for _, book := range l.Books {
		fmt.Println(book.DisplayDetails())
	}
}

// Main function to demonstrate enhanced library management
func main() {
	lib := Library{}

	// Add books and ebooks
	book1 := NewBook("Go Programming Basics", "John Doe", "123456789", true)
	ebook1 := EBook{
		Book:     book1,
		FileSize: 5,
	}

	lib.AddBook(book1)
	lib.AddBook(ebook1)

	// List books and ebooks
	fmt.Println("Books and EBooks in the Library:")
	lib.ListBooks()
}
