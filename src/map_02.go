package main

import "fmt"

func main() {
	personSalary := map[string]int{
		"aa": 20000,
		"bb": 30000,
	}
	personSalary["cc"] = 9000
	fmt.Println("all items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
}
