package mylibrary

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

var Books = []Book{{Id: 1, Title: "Book1", Author: "Author1"}, {Id: 2, Title: "Book2", Author: "Author2"}}

func BookGet(g *gin.RouterGroup) {
	g.GET("/books", func(c *gin.Context) {
		c.JSON(200, Books)
	})
}

func BookPost(g *gin.RouterGroup) {
	g.POST("/books", func(c *gin.Context) {
		var book Book
		c.BindJSON(&book)
		// add ID
		book.Id = Books[len(Books)-1].Id + 1
		Books = append(Books, book)
		c.JSON(200, book)
	})
}

func BookPut(g *gin.RouterGroup) {
	g.PUT("/books/:id", func(c *gin.Context) {
		var book Book
		c.BindJSON(&book)
		// add ID
		id, _ := strconv.Atoi(c.Param("id"))
		book.Id = id
		Books = append(Books, book)
		c.JSON(200, book)
	})
}

func BookDelete(g *gin.RouterGroup) {
	g.DELETE("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i, book := range Books {
			if book.Id == id {
				Books = append(Books[:i], Books[i+1:]...)
				break
			}
		}
		c.JSON(200, Books)
	})
}
