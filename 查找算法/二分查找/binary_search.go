package main

import (
	"fmt"
)

func main() {
	item := 20
	list := []int{
		9, 10, 20, 50,
	}

	index,count := binary_search(list, item)
	fmt.Printf("Result\n    - index: %d\n    - count: %d\n", index, count)
}

func binary_search(list []int, item int) (int, int) {
	low := 0
	high := len(list) - 1

	count := 0
	for low <= high {
		count++

		mid := (low + high) / 2
		guess := list[mid]

		fmt.Printf("%d %d %d %d\n", low, high, mid, guess)

		if guess == item {
			return mid,count
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1,count
}