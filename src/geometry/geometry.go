package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var aa, bb float64 = 7, 9

func init() {
	println("use package initialized")
	if aa < 0 {
		log.Fatal("lenth is less than 0")
	}
	if bb < 0 {
		log.Fatal("width is less than 0")
	}
}
func main() {
	fmt.Println("Begin")

	fmt.Printf("area is %.2f\n", rectangle.Area(aa, bb))

	fmt.Printf("diagonal is %2.f\n", rectangle.Diagonal(aa, bb))
}
