package main

import (
	"fmt"
	"playground/internal/scalability"
)

func main() {
	var nums []int = []int{1, 7, 4, 9, 3, 6, 2}
	fmt.Println(scalability.BubbleSort(nums))
}
