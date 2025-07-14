# 📚 Console-Based Library Management System

## 🎯 Objective
Create a simple console-based library management system in Go using:
- Structs
- Interfaces
- Methods
- Maps
- Slices

---

## 🧱 Structs

### `Book`
- `ID` (int)
- `Title` (string)
- `Author` (string)
- `Status` (string) — _"Available"_ or _"Borrowed"_

### `Member`
- `ID` (int)
- `Name` (string)
- `BorrowedBooks` (`[]Book`)

---

## 🧩 Interface: `LibraryManager`

- `AddBook(book Book)`
- `RemoveBook(bookID int)`
- `BorrowBook(bookID int, memberID int) error`
- `ReturnBook(bookID int, memberID int) error`
- `ListAvailableBooks() []Book`
- `ListBorrowedBooks(memberID int) []Book`

---

## 🏗️ Implementation Details

### `Library` Struct
Implements the `LibraryManager` interface.

#### Fields:
- `Books`: `map[int]Book`
- `Members`: `map[int]Member`

---

## ✅ Features Implemented

1. **AddBook** – Adds a new book to the collection.
2. **RemoveBook** – Deletes a book by its ID.
3. **BorrowBook** – A member borrows a book if it's available.
4. **ReturnBook** – A member returns a borrowed book.
5. **ListAvailableBooks** – Prints all available books.
6. **ListBorrowedBooks** – Lists books borrowed by a specific member.
7. **AddMember** – Adds a new member to the library.

---

## 📂 Folder Structure

library_management/
├── main.go // Entry point
├── controllers/
│ └── library_controller.go // Handles user input
├── models/
│ ├── book.go // Book struct
│ └── member.go // Member struct
├── services/
│ └── library_service.go // Logic & data operations
├── docs/
│ └── documentation.md // ← This file
└── go.mod // Go module file


---

## 📝 Evaluation Criteria

- ✅ Correct implementation of all methods
- ✅ Functional input/output via console
- ✅ Proper error handling for missing books/members
- ✅ Organized folder and file structure
- ✅ Clean, readable, and maintainable code

---

## 🗒️ Notes

- All input/output is handled via the **console**.
- Data is stored **in memory** (no file/database persistence).
- Ideal project for **practicing Go fundamentals**.

---
