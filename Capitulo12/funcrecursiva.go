package main

import "fmt"


func main() {

	x := factorial(5)
	fmt.Println(x)

}

func factorial( x int) int {
	if x == 1 {
		return x
	} 
	return x * factorial(x - 1)
}