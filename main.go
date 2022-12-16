package main

import (
	"sample/mylibrary"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	// create group /api/v1
	g := r.Group("/api/v1")

	mylibrary.UserGet(g)
	mylibrary.UserPost(g)
	mylibrary.UserPut(g)
	mylibrary.UserDelete(g)
	mylibrary.BookGet(g)
	mylibrary.BookPost(g)
	mylibrary.BookPut(g)
	mylibrary.BookDelete(g)

	r.Run(":8080")
}
