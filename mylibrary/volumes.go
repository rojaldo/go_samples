package mylibrary

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Volumes = []Volume{{Id: 1, BookId: 1, damages: ""}, {Id: 2, BookId: 1, damages: "damages2"}, {Id: 3, BookId: 2, damages: "damages3"}}

type Volume struct {
	Id      int
	BookId  int
	damages string
}

func VolumeGet(g *gin.RouterGroup) {
	g.GET("/volumes", func(c *gin.Context) {
		c.JSON(200, Volumes)
	})
}

func VolumePost(g *gin.RouterGroup) {
	g.POST("/volumes", func(c *gin.Context) {
		var volume Volume
		c.BindJSON(&volume)
		// check that BookId is in Books
		// if not, return error
		for _, book := range Books {
			if book.Id == volume.BookId {
				// check that BookId is not in transactions
				for _, transaction := range Transactions {
					if transaction.BookVolumeId == volume.BookId {
						c.JSON(400, errors.New("BookId is already in transactions"))
						return
					}
				}
				// add ID
				volume.Id = Volumes[len(Volumes)-1].Id + 1
				Volumes = append(Volumes, volume)
				c.JSON(201, volume)
				return
			}
		}
		c.JSON(400, errors.New("BookId not found"))
	})
}

func VolumePut(g *gin.RouterGroup) {
	g.PUT("/volumes/:id", func(c *gin.Context) {
		var volume Volume
		c.BindJSON(&volume)
		// check that BookId is in Books
		// if not, return error
		for _, book := range Books {
			if book.Id == volume.BookId {
				// add ID
				id, _ := strconv.Atoi(c.Param("id"))
				volume.Id = id
				Volumes = append(Volumes, volume)
				c.JSON(200, volume)
				return
			}
		}
		c.JSON(400, errors.New("BookId not found"))
	})
}

func VolumeDelete(g *gin.RouterGroup) {
	g.DELETE("/volumes/:id", func(c *gin.Context) {
		id, error := strconv.Atoi(c.Param("id"))
		if error != nil {
			c.JSON(400, errors.New("Not a valid VolumeId"))
			return
		}
		for i, volume := range Volumes {
			if volume.Id == id {
				Volumes = append(Volumes[:i], Volumes[i+1:]...)
				c.JSON(200, volume)
				return
			}
		}
		c.JSON(400, errors.New("VolumeId not found"))
	})
}
