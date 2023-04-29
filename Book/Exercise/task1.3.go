package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	Lastname string
	Number   int
}

func binarySearchName(arr []Person, item string) (int, time.Duration) {
	start := time.Now()
	low := 0
	mid := 0
	guess := ""
	high := len(arr) - 1
	for low <= high {
		mid = (high + low) / 2
		guess = arr[mid].Lastname
		if guess == item {
			duration := time.Since(start)
			return arr[mid].Number, duration
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, 0

}

func SearchName(list []Person, item string) (int, time.Duration) {
	start := time.Now()
	for i := 0; i < len(list); i++ {
		if list[i].Lastname == item {
			duration := time.Since(start)
			return list[i].Number, duration
		}
	}
	return 0, 0
}

func main() {

	list := []Person{
		{"Oleg", 98988},
		{"Gleb", 97743},
		{"Lena", 38498},
		{"Viktor", 3748},
		{"Mika", 87434},
		{"Rik", 98988},
		{"Anna", 97743},
		{"Nasty", 38498},
		{"Roma", 3748},
		{"Masha", 87434},
	}
	sort.Slice(list, func(i, j int) (less bool) {
		return list[i].Lastname < list[j].Lastname
	})

	fmt.Println(list)

	fmt.Print("Binary search O(log n) ")
	fmt.Println(binarySearchName(list, "Masha"))

	fmt.Print("Linear search O(n) ")
	fmt.Println(SearchName(list, "Masha"))

}
