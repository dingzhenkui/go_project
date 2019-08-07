package main

import "fmt"

func main() {
	personSalary := map[string]int{
		"steve": 12000,
		"kamie": 15000,
	}
	personSalary["mike"] = 9000
	newEmp := "mike"
	value, ding := personSalary[newEmp]
	if ding == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}
}
