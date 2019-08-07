package main

import "fmt"

func main() {
	em3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Tom",
		lastName:  "ss",
		age:       26,
		salary:    803,
	}
	fmt.Println("Employee 3", em3)
}
