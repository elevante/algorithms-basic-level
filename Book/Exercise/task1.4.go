package main

import (
	"fmt"
	"sort"
	"time"
)

func SearchNumber(list []Person, item int) (string, time.Duration) {
	start := time.Now()
	for i := 0; i < len(list); i++ {
		if list[i].Number == item {
			duration := time.Since(start)
			return list[i].Lastname, duration
		}
	}
	return "", 0
}

func main() {

	list := []Person{
		{"Oleg", 98978},
		{"Gleb", 97743},
		{"Lena", 38498},
		{"Viktor", 3748},
		{"Mika", 87434},
		{"Rik", 98983},
		{"Anna", 97743},
		{"Nasty", 38498},
		{"Roma", 3748},
		{"Masha", 87434},
	}
	sort.Slice(list, func(i, j int) (less bool) {
		return list[i].Number < list[j].Number
	})

	fmt.Println(list)

	fmt.Print("Linear search O(n) ")
	fmt.Println(SearchNumber(list, 98983))

}
