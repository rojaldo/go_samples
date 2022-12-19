package mylibrary

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func startServer() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	g := r.Group("/api/v1")
	UserGet(g)
	r.Run(":8080")
}

func TestMain(t *testing.T) {
	// Setup
	go startServer()
	client := resty.New()

	// Execute
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:8080/api/v1/users")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 200, resp.StatusCode(), "they should be equal")

	// Execute
	resp, err = client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:8080/api/v1/users/1")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 200, resp.StatusCode(), "they should be equal")

	// Execute
	resp, err = client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:8080/api/v1/users/23")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 404, resp.StatusCode(), "they should be equal")
}
