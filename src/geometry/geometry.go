package main

import (
	"fmt"
	"geometry/rectangle"
)

func main() {
	var aa, bb float64 = 7, 9
	fmt.Println("Begin")

	fmt.Printf("area is %.2f\n", rectangle.Area(aa, bb))

	fmt.Printf("diagonal is %2.f\n", rectangle.Diagonal(aa, bb))
}
