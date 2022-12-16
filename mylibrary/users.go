package mylibrary

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int
	Username string
	Password string
}

var Users = []User{{Id: 1, Username: "user1", Password: "pass1"}, {Id: 2, Username: "user2", Password: "pass2"}}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users [get]
func UserGet(g *gin.RouterGroup) {
	g.GET("/users", func(c *gin.Context) {
		c.JSON(200, Users)
	})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users [post]
func UserPost(g *gin.RouterGroup) {
	g.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		// add ID
		user.Id = Users[len(Users)-1].Id + 1
		Users = append(Users, user)
		c.JSON(200, user)
	})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users [put]
func UserPut(g *gin.RouterGroup) {
	g.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// get int from string
		idInt, error := strconv.Atoi(id)
		if error != nil {
			c.JSON(400, gin.H{"error": error})
		}
		var user User
		c.BindJSON(&user)
		for i, u := range Users {
			if u.Id == idInt {
				Users[i] = user
				break
			}
		}

	})
}

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users [delete]
func UserDelete(g *gin.RouterGroup) {
	g.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, error := strconv.Atoi(id)
		if error != nil {
			c.JSON(400, gin.H{"error": error})
		}
		for i, u := range Users {
			if u.Id == idInt {
				Users = append(Users[:i], Users[i+1:]...)
				break
			}
		}
	})
}
