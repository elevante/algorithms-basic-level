package main

import "fmt"

func binarySearch(arr []int, item int) (int, int) {
	low := 0
	mid := 0
	guess := 0
	high := len(arr) - 1
	count := 0
	for low <= high {
		mid = (low + high) / 2
		guess = arr[mid]

		if guess == item {
			return mid, count
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
		count++
	}
	return 0, 0
}

func main() {
	var arr []int
	var arr1 []int
	for a := 1; a <= 128; a++ {
		arr = append(arr, a)
	}

	fmt.Println(binarySearch(arr, 128))

	for a := 0; a <= 256; a++ {
		arr1 = append(arr1, a)
	}

	println(binarySearch(arr1, 256))

}
