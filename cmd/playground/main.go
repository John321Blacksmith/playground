package main

import (
	"fmt"
	// "playground/internal/scalability"
)

func PutItem(item int, rec chan<- int) {
	rec <- item
}

func main() {
	intChan := make(chan int)
	for i := range 10 {
		go PutItem(i, intChan)
	}
	fmt.Println(<-intChan)
}

// fork -join model
