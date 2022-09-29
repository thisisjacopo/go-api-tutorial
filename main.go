package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID		string	`json: "id"`
	Title 	string	`json: "title"`
	Author	string	`json: "author"`
	Quantity int	`json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "Title number 1", Author: "Author number 1", Quantity: 2}
	{ID: "2", Title: "Title number 2", Author: "Author number 2", Quantity: 21}
	{ID: "3", Title: "Title number 3", Author: "Author number 3", Quantity: 10}
	{ID: "4", Title: "Title number 4", Author: "Author number 4", Quantity: 54}
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func main() {
	router := gin.Default()
	rounter.GET("/books", getBooks)
	rounter.Run("localhost:8080")
}