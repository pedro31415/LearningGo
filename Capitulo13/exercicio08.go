package main

import "fmt"

func main () {
	x := sum()

	fmt.Println(x(5))
}


func sum() func(y int) int {
	x := 10
	return func(y int) int {
		return x + y
	}
} 