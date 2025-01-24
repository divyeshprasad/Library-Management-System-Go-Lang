package main

import (
	"fmt"
)

type Book struct {
	Title     string
	Author    string
	Isbn      string
	Available bool
}

func NewBook(title string, author string, isbn string, available bool) Book {
	return Book{Title: title, Author: author, Isbn: isbn, Available: available}
}

func (b Book) DisplayDetails() {
	fmt.Printf("Title:%s\n Author:%s\n Isbn:%s\n Available:%t\n", b.Title, b.Author, b.Isbn, b.Available)
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(isbn string) error {
	for i, book := range l.Books {
		if book.Isbn == isbn {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with ISBN %s not found", isbn)
}

func (l *Library) SearchBookByTitle(title string) []Book {
	var result []Book
	for _, book := range l.Books {
		if contains(book.Title, title) {
			result = append(result, book)
		}
	}
	return result
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}

func (l *Library) ListBooks() {
	// if len(l.Books) == 0 {
	//     fmt.Println("No books available.")
	//     return
	// }
	for _, book := range l.Books {

		book.DisplayDetails()
	}
}

func main() {
	lib := Library{}

	// Add books
	book1 := NewBook("Go Programming Basics", "John Doe", "123456789", true)
	book2 := NewBook("Advanced Go Programming", "Jane Smith", "987654321", true)

	lib.AddBook(book1)
	lib.AddBook(book2)

	// List books
	fmt.Println("Books in the Library:")
	lib.ListBooks()

	// Remove a book
	fmt.Println("\nRemoving book with ISBN 123456789...")
	err := lib.RemoveBook("123456789")
	if err != nil {
		fmt.Println(err)
	}

	// List books again
	fmt.Println("\nBooks in the Library after removal:")
	lib.ListBooks()

	// Search books by title
	fmt.Println("\nSearch for books with 'Go' in the title:")
	res := lib.SearchBookByTitle("Go")
	for _, book := range res {
		book.DisplayDetails()
	}
}
