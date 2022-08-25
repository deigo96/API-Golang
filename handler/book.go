package handler

import (
	"errors"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, book.Books)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CheckOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing query parameter"})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookTitle(c *gin.Context) {
	title := c.Query("title")

	book, err := GetBookByTitle(title)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": ok})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*book.Book, error) {
	for i, b := range book.Books {
		if b.ID == id {
			return &book.Books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func GetBookByTitle(title string) (*book.Book, error) {
	for i, b := range book.Books {
		if b.Title == title {
			return &book.Books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func CreateBook(c *gin.Context) {
	var newBook book.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	book.Books = append(book.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}