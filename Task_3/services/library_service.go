package services

import (
	"fmt"

	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book

}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func (l *Library) AddBook(book models.Book) {
	if l.Books == nil {
		l.Books = make(map[int]models.Book)
	}
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(BookId int) {
	delete(l.Books, BookId) // delete(map, key)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	if book.Status == "Borrowed" {
		return fmt.Errorf("book is already borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    book, bookExists := l.Books[bookID]
    if !bookExists {
        return fmt.Errorf("book with ID %d not found", bookID)
    }

    member, memberExists := l.Members[memberID]
    if !memberExists {
        return fmt.Errorf("member with ID %d not found", memberID)
    }

    // Check if the member actually borrowed this book
    found := false
    updatedBorrowedBooks := []models.Book{}
    for _, b := range member.BorrowedBooks {
        if b.ID == bookID {
            found = true
            continue // skip the returned book
        }
        updatedBorrowedBooks = append(updatedBorrowedBooks, b)
    }

    if !found {
        return fmt.Errorf("member did not borrow this book")
    }

    // Update member's borrowed list and book status
    member.BorrowedBooks = updatedBorrowedBooks
    l.Members[memberID] = member
    book.Status = "Available"
    l.Books[bookID] = book

    return nil
}

func (l *Library) ListAvailableBooks()  []models.Book{
	// return available books
	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
