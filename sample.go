package main

import "fmt"

const (
	Summer = "summer"
	Autumn = "autumn"
	Winter = "winter"
	Spring = "spring"
)

const (
	Apples  = 0
	Oranges = 1
)

func main() {
	// Ideally, this should never be true!
	fmt.Println(Summer == Apples)
}
