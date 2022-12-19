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
	UserPost(g)
	UserPut(g)
	UserDelete(g)
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

	// setup
	Users = []User{{Id: 1, Username: "user1", Password: "pass1"}, {Id: 2, Username: "user2", Password: "pass2"}}

	// Execute post request with body
	resp, err = client.R().
		SetHeader("Accept", "application/json").
		SetBody(`{"username": "user3", "password": "pass3"}`).
		Post("http://localhost:8080/api/v1/users")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 201, resp.StatusCode(), "they should be equal")
	assert.Equal(t, 3, len(Users), "they should be equal")
	assert.Equal(t, `{"id":3,"username":"user3","password":"pass3"}`, resp.String(), "they should be equal")

	// Teardown
	Users = []User{{Id: 1, Username: "user1", Password: "pass1"}, {Id: 2, Username: "user2", Password: "pass2"}}

	runPutTest(t, client)
	// Execute delete request with body

	resp, err = client.R().
		SetHeader("Accept", "application/json").
		Delete("http://localhost:8080/api/v1/users/1")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 200, resp.StatusCode(), "they should be equal")
	assert.Equal(t, 1, len(Users), "they should be equal")
	assert.Equal(t, `{"message":"user deleted"}`, resp.String(), "they should be equal")

}

func runPutTest(t *testing.T, client *resty.Client) {
	// Execute put request with body

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(`{"username": "user3", "password": "pass3"}`).
		Put("http://localhost:8080/api/v1/users/1")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 200, resp.StatusCode(), "they should be equal")
	assert.Equal(t, 2, len(Users), "they should be equal")
	assert.Equal(t, `{"id":1,"username":"user3","password":"pass3"}`, resp.String(), "they should be equal")

	// get user 1
	resp, err = client.R().
		SetHeader("Accept", "application/json").
		Get("http://localhost:8080/api/v1/users/1")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 200, resp.StatusCode(), "they should be equal")
	assert.Equal(t, `{"id":1,"username":"user3","password":"pass3"}`, resp.String(), "they should be equal")

	// Execute a put request with body of invalid ID

	resp, err = client.R().
		SetHeader("Accept", "application/json").
		SetBody(`{"username": "user3", "password": "pass3"}`).
		Put("http://localhost:8080/api/v1/users/23")
	if err != nil {
		t.Error("Error")
	}

	// assert
	assert.Equal(t, 404, resp.StatusCode(), "they should be equal")

	// Tear down
	Users = []User{{Id: 1, Username: "user1", Password: "pass1"}, {Id: 2, Username: "user2", Password: "pass2"}}
}
