package main

import "fmt"

func main () {
	const x = 123
	const valueNumber float64 = 123.567

	fmt.Printf("strong typing -> %.3f\n", valueNumber)
	fmt.Print("weak typing -> %d", x)
}