package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)


/*AddBook(book Book)
RemoveBook(bookID int)
BorrowBook(bookID int, memberID int) error
ReturnBook(bookID int, memberID int) error
ListAvailableBooks() []Book
ListBorrowedBooks(memberID int) []Book
*/
func main() {
	// Initialize the library
	library := &services.Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			controllers.AddBookController(library)
		case 2:
			controllers.RemoveBookController(library)
		case 3:
			controllers.BorrowBookController(library)
		case 4:
			controllers.ReturnBookController(library)
		case 5:
			controllers.ListAvailableBooksController(library)
		case 6:
			controllers.ListBorrowedBooksController(library)
		case 0:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
