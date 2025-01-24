# Library Management System in Go

This project implements a simple **Library Management System** in Go, which allows users to manage a collection of books and eBooks. The system provides functionalities to add books, remove books by ISBN, search for books by title, and list all available books.

## Project Structure

- **part1.go**: Implements the basic library management system with operations such as adding, removing, and listing books.
- **part2.go**: Enhances the system by adding support for eBooks, using inheritance through embedding and implementing interfaces.
- **part3.go**: Implements a text-based menu system allowing users to interact with the library system.

## Requirements

- Go 1.18+ (or later)

## How to Run the Program

### Prerequisites

Ensure that you have Go installed. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

### Running the Program

1. Download or clone the repository to your local machine.
2. Navigate to the project directory in your terminal.

#### To run **part1.go**:
```bash
go run part1.go
```

#### To run **part2.go**:
```bash
go run part2.go
```

#### To run **part3.go**:
```bash
go run part3.go
```

You can test each part of the system by running the corresponding Go file.

## Features

### **Part 1: Basic Library Management System**

- Add books to the library.
- Remove books by ISBN.
- Search for books by title.
- List all books in the library.

### **Part 2: Enhanced Library Management System**

- Adds support for eBooks, which embed the `Book` struct and include additional information (file size).
- Implements a `BookInterface` to allow polymorphic behavior (both `Book` and `EBook` conform to this interface).
  
### **Part 3: Text-Based Menu System**

- Interactive menu for adding, removing, and listing books and eBooks.
- Input is taken from the user via the terminal.

---

## Example of Input/Output

### 1. **Adding a Book**

When you run the program and choose the option to **Add a Book/EBook**, the system will ask you to enter the details of the book or eBook.

**Example Input:**
```
Enter title: Go Programming Basics
Enter author: John Doe
Enter ISBN: 123456789
Is available (true/false): true
Enter file size (in MB for EBook, 0 for Book): 0
```

**Example Output:**
```
Book added successfully.
```

### 2. **Adding an EBook**

For adding an eBook, you need to enter a non-zero value for the file size.

**Example Input:**
```
Enter title: Advanced Go Programming
Enter author: Jane Smith
Enter ISBN: 987654321
Is available (true/false): true
Enter file size (in MB for EBook, 0 for Book): 5
```

**Example Output:**
```
EBook added successfully.
```

### 3. **Listing Books and EBooks**

To list all books and eBooks in the library, select the **List all Books/EBooks** option.

**Example Output:**
```
Title: Go Programming Basics, Author: John Doe, ISBN: 123456789, Available: true
Title: Advanced Go Programming, Author: Jane Smith, ISBN: 987654321, Available: true, FileSize: 5MB
```

### 4. **Removing a Book by ISBN**

To remove a book, select the **Remove a Book/EBook** option and enter the ISBN of the book you wish to remove.

**Example Input:**
```
Enter ISBN of the book you want to remove: 123456789
```

**Example Output:**
```
Book removed successfully.
```

### 5. **Handling Errors**

If you try to remove a book that does not exist in the library, an error will be shown.

**Example Input:**
```
Enter ISBN of the book you want to remove: 999999999
```

**Example Output:**
```
Error: book with ISBN 999999999 not found
```

---

## Notes

- The program uses slices to manage collections of books and eBooks.
- The `Book` struct represents a basic book, while the `EBook` struct embeds `Book` and adds additional information such as file size.
- The program supports polymorphism through the `BookInterface`, which both `Book` and `EBook` types implement.
- Basic error handling is included, such as when an ISBN is not found when removing a book.

## ScreenShots
![Part1]{Screenshots/Part1.png}
![Part2]{Screenshots/Part2.png}
![Part3]{Screenshots/Part3.png}

