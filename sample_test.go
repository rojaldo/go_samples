package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	response := helloWorld()

	assert.Equal(t, "Hello, world", response, "they should be equal")

}
