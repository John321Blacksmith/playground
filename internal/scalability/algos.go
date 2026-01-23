package scalability

import (
	"fmt"
)

// time complexity is O(n*2)
func BubbleSort(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
			fmt.Println(arr)

		}
	}
	return arr
}

// Time complexity O(nlogn)
func QuickSort(arr []int) []int {

	// If the array has been divided to limits...
	if len(arr) < 1 {
		return arr
	}

	// Specify boundaries
	left, right := 0, len(arr)-1

	// Pick a pivot and move it to the end
	pivot := (len(arr) - 1) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Partition
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Putting the pivot in place
	arr[left], arr[right] = arr[right], arr[left]

	// Perform QuickSort recursively with sub-arrays
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
