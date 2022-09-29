package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID		 string	`json:"id"`
	Title 	 string	`json:"title"`
	Author	 string	`json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Title number 1", Author: "Author number 1", Quantity: 2},
	{ID: "2", Title: "Title number 2", Author: "Author number 2", Quantity: 21},
	{ID: "3", Title: "Title number 3", Author: "Author number 3", Quantity: 10},
	{ID: "4", Title: "Title number 4", Author: "Author number 4", Quantity: 54},
}


func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
	
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}