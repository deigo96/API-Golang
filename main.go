package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// _ untuk variabel yang tidak dipake
	// koneksi database
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error") // log.Fatal() untuk menstop go
	}

	db.AutoMigrate(&book.Buku{}) // auto migrate membuat tabel di database

	// bookRepository := book.NewRepository(db)

	// book := book.Buku{
	// 	Title: "Macan Lapar",
	// 	Description: "Macan yang sedang kelaparan mencari makan",
	// 	Price: 91000,
	// 	Rating: 5,
	// }

	router := gin.Default()

	// api versioning
	v1 := router.Group("/v1")
	v1.GET("/books", handler.GetBooks)
	v1.GET("/books/:id", handler.BookById)
	v1.GET("/books/title", handler.GetBookTitle)
	v1.GET("/rapidAPI/title", handler.AutoComplete)
	v1.GET("/rapidAPI/id", handler.DetailMovie)
	v1.POST("/books", handler.CreateBook)
	v1.PATCH("/checkout", handler.CheckOutBook)
	v1.PATCH("/adding", handler.ReturnBook)
	router.Run("localhost:9999")
}
