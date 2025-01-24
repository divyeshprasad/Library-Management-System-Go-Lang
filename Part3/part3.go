package main

import (
	"fmt"
)

// Book struct with basic details
type Book struct {
	Title    string
	Author   string
	ISBN     string
	Available bool
}

// Method to initialize a new book
func NewBook(title, author, isbn string, available bool) Book {
	return Book{
		Title:    title,
		Author:   author,
		ISBN:     isbn,
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

// Library struct to manage collection of BookInterface (Book and EBook)
type Library struct {
	Books []interface{}
}

// Add a book or ebook to the library
func (l *Library) AddBook(book interface{}) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(isbn string) error {
	for i, book := range l.Books {
		switch b := book.(type) {
		case Book:
			if b.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				return nil
			}
		case EBook:
			if b.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				return nil
			}
		}
	}
	return fmt.Errorf("book with ISBN %s not found", isbn)
}

// List all books and ebooks in the library
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books available.")
		return
	}
	for _, book := range l.Books {
		switch b := book.(type) {
		case Book:
			fmt.Println(b.DisplayDetails())
		case EBook:
			fmt.Println(b.DisplayDetails())
		}
	}
}

func main() {
	lib := Library{}
	var choice int

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book/EBook")
		fmt.Println("2. Remove a Book/EBook")
		fmt.Println("3. List all Books/EBooks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title, author, isbn string
			var available bool
			var fileSize int
			fmt.Print("Enter title: ")
			fmt.Scan(&title)
			fmt.Print("Enter author: ")
			fmt.Scan(&author)
			fmt.Print("Enter ISBN: ")
			fmt.Scan(&isbn)
			fmt.Print("Is available (true/false): ")
			fmt.Scan(&available)
			fmt.Print("Enter file size (in MB for EBook, 0 for Book): ")
			fmt.Scan(&fileSize)

			if fileSize > 0 {
				ebook := EBook{
					Book:     Book{Title: title, Author: author, ISBN: isbn, Available: available},
					FileSize: fileSize,
				}
				lib.AddBook(ebook)
			} else {
				book := Book{Title: title, Author: author, ISBN: isbn, Available: available}
				lib.AddBook(book)
			}
		case 2:
			var isbnToRemove string
			fmt.Print("Enter ISBN of the book you want to remove: ")
			fmt.Scan(&isbnToRemove)
			err := lib.RemoveBook(isbnToRemove)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book removed successfully.")
			}
		case 3:
			lib.ListBooks()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
