// module documentation for the Library Management System

module library_management_documentation

go 1.21

// ===========================================
// Console-Based Library Management System
// ===========================================
// Objective:
// Create a simple console-based library management system in Go 
// using structs, interfaces, methods, maps, and slices.

// ===========================================
// Structs
// ===========================================
// Book struct:
// - ID (int)
// - Title (string)
// - Author (string)
// - Status (string) // "Available" or "Borrowed"

// Member struct:
// - ID (int)
// - Name (string)
// - BorrowedBooks ([]Book)

// ===========================================
// Interface: LibraryManager
// ===========================================
// - AddBook(book Book)
// - RemoveBook(bookID int)
// - BorrowBook(bookID int, memberID int) error
// - ReturnBook(bookID int, memberID int) error
// - ListAvailableBooks() []Book
// - ListBorrowedBooks(memberID int) []Book

// ===========================================
// Implementation Details
// ===========================================
// Library struct:
// - Implements LibraryManager interface
// - Fields:
//   - Books: map[int]Book
//   - Members: map[int]Member

// ===========================================
// Features Implemented
// ===========================================
// 1. AddBook: Adds a new book to the collection.
// 2. RemoveBook: Deletes a book by its ID.
// 3. BorrowBook: A member borrows a book if available.
// 4. ReturnBook: A member returns a borrowed book.
// 5. ListAvailableBooks: Prints all available books.
// 6. ListBorrowedBooks: Lists books borrowed by a member.
// 7. AddMember: Adds a new member to the library.

// ===========================================
// Folder Structure
// ===========================================
// library_management/
// ├── main.go                          // entry point
// ├── controllers/
// │   └── library_controller.go       // handles user input
// ├── models/
// │   ├── book.go                     // Book struct
// │   └── member.go                   // Member struct
// ├── services/
// │   └── library_service.go          // logic & data operations
// ├── docs/
// │   └── documentation.mod           // ← this file
// └── go.mod                          // Go module file

// ===========================================
// Evaluation Criteria
// ===========================================
// ✅ Correct implementation of all methods
// ✅ Functional input/output via console
// ✅ Proper error handling for missing books/members
// ✅ Organized folder and file structure
// ✅ Written with clean, maintainable code

// ===========================================
// Notes
// ===========================================
// - All input/output is handled via the console.
// - Data is stored in memory (not persistent).
// - This project is ideal for practicing Go fundamentals.
