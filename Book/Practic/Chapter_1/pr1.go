package main

func binary_search(arr []int, item int) int {
	low := 0
	mid := 0
	var guess int

	high := len(arr) - 1

	for low <= high {
		mid = (low + high) / 2
		guess = arr[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

func pr() {
	//employee := new(Employee)
	//employee.fullName = "fjkdjfkjdf"
	//fmt.Println(employee.fullName)
}
