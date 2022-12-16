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
	mylibrary.VolumeGet(g)
	mylibrary.VolumePost(g)
	mylibrary.VolumePut(g)
	mylibrary.VolumeDelete(g)
	mylibrary.TransactionGet(g)
	mylibrary.TransactionPost(g)
	mylibrary.TransactionDelete(g)

	// angular app with multiple routes
	r.Static("/app", "./app")
	r.LoadHTMLGlob("app/*.html")
	// allow access to all files in app folder

	r.StaticFile("/app", "./*.*")

	r.Run(":8080")
}
