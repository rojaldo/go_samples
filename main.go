package main

import "github.com/gin-gonic/gin"

type User struct {
	Username string
	Password string
}

var users = []User{{Username: "user1", Password: "pass1"}, {Username: "user2", Password: "pass2"}}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		users = append(users, user)
		c.JSON(200, user)
	})

	r.PUT("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		var user User
		c.BindJSON(&user)
		for i, u := range users {
			if u.Username == username {
				users[i] = user
				break
			}
		}
		c.JSON(200, user)
	})

	r.DELETE("/users/:username", func(c *gin.Context) {
		username := c.Param("username")
		for i, u := range users {
			if u.Username == username {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		c.JSON(200, gin.H{"username " + username: "deleted"})
	})

	r.Run(":8080")
}
