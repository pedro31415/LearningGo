package main

import "fmt"

func main() {
	value := sum(30, 20)

	fmt.Println(value)
}

func sum(x ...int) int {
	result := 0 

	for _, total := range x {
		result += total
	}

	return result
}