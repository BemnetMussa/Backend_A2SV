package controllers

import (
    "bufio"
    "fmt"
    "../models"
    "../services"
    "os"
    "strconv"
    "strings"
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
