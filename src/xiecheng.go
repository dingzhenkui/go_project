package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func alphabets() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	go number()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("aaaaaaaaa")
}
