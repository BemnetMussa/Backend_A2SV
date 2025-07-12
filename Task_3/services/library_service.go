package services

import (
	"../models"
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
	Books map[int]models.Book
	members map[int]models.Member
}

func (l *Library) AddBook(book models.Book) {
	if l.Books == nil {
		l.Books = make(map[int]models.Book)
	}
	l.Books[book.ID] = book
}