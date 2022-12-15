package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	getTetrisPiece()
}

func getTetrisPiece() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(7))

}
