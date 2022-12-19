package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	// Create an observable that emits a series of integers
	observable := rxgo.Just("Hello, World!")()
	ch := observable.Observe()
	item := <-ch
	fmt.Println(item.V)

	// Create an observable that emits a series of integers
	observable = rxgo.Just(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)()
	ch = observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}

	// Create an observable that gets the response from a REST API
	observable = rxgo.Just(getResponse())()
	ch = observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
}

// getResponse returns a channel that emits the response from a REST API
func getResponse() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		ch <- "Hello, World!"
		close(ch)
	}()
	return ch
}
