package mylibrary

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Transactions = []Transaction{{Id: 1, BookVolumeId: 1, UserId: 1, TransactionDate: "2020-01-01", DueDate: "2020-01-15"}}

type Transaction struct {
	Id              int
	BookVolumeId    int
	UserId          int
	TransactionDate string
	DueDate         string
}

func TransactionGet(g *gin.RouterGroup) {
	g.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, Transactions)
	})
}

func TransactionPost(g *gin.RouterGroup) {
	g.POST("/transactions", func(c *gin.Context) {
		var transaction Transaction
		c.BindJSON(&transaction)
		// check that BookVolumeId is in Volumes
		// if not, return error
		for _, volume := range Volumes {
			if volume.Id == transaction.BookVolumeId {
				// check that UserId is in Users
				// if not, return error
				for _, user := range Users {
					if user.Id == transaction.UserId {
						// add ID
						transaction.Id = Transactions[len(Transactions)-1].Id + 1
						Transactions = append(Transactions, transaction)
						c.JSON(201, transaction)
						return
					}
				}
				c.JSON(400, errors.New("UserId not found"))
				return
			}
		}
		c.JSON(400, errors.New("BookVolumeId not found"))
	})
}

func TransactionDelete(g *gin.RouterGroup) {
	g.DELETE("/transactions/:id", func(c *gin.Context) {
		id, error := strconv.Atoi(c.Param("id"))
		if error != nil {
			c.JSON(400, errors.New("Not a valid TransactionId"))
			return
		}
		for i, transaction := range Transactions {
			if transaction.Id == id {
				Transactions = append(Transactions[:i], Transactions[i+1:]...)
				c.JSON(200, transaction)
				return
			}
		}
		c.JSON(400, errors.New("TransactionId not found"))
	})
}
