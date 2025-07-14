package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

func AddBookController(manager services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Book ID: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Author Name: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	manager.AddBook(book)
	fmt.Println("Book added successfully!")
}

func RemoveBookController(manager services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Book ID: ")
	idInput, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idInput))

	manager.RemoveBook(id)
}

func BorrowBookController(manager services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Book ID: ")
	input, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Enter Member ID: ")
	input, _ = reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(input))

	err := manager.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Borrow failed:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func ReturnBookController(manager services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Book ID: ")
	input, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(input))

	fmt.Print("Enter Member ID: ")
	input, _ = reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(input))

	err := manager.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Borrow failed: ", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func ListAvailableBooksController(manager services.LibraryManager) {
	availableBooks := manager.ListAvailableBooks()

	if len(availableBooks) == 0 {
		fmt.Println("No available books.")
		return
	}

	for _, book := range availableBooks {
		fmt.Printf("Book ID: %d\nTitle: %s\nAuthor: %s\nStatus: %s\n\n",
			book.ID, book.Title, book.Author, book.Status)
	}
}

func ListBorrowedBooksController(manager services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Member ID: ")
	input, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid member ID.")
		return
	}

	borrowedBooks := manager.ListBorrowedBooks(memberID)

	if borrowedBooks == nil || len(borrowedBooks) == 0 {
		fmt.Println("No borrowed books found for this member.")
		return
	}

	for _, book := range borrowedBooks {
		fmt.Printf("Book ID: %d\nTitle: %s\nAuthor: %s\nStatus: %s\n\n",
			book.ID, book.Title, book.Author, book.Status)
	}
}
