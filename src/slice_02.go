package main

import "fmt"

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dbbb := darr[2:5]
	for i := range dbbb {
		dbbb[i]++
	}
	fmt.Println("dsds", darr)
}
