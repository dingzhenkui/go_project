package main

import "fmt"

func main() {
	a := [...]float64{65.4, 54, 76.5, 32}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d dsdfs %.2f\n", i, v)
		sum += v
	}
	fmt.Println("sum is ", sum)
}
