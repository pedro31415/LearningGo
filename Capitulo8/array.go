package main

import "fmt"

var x [5]int

func main () {
	for i:=0; i <5; i++ {
		x[i] = 10 * i;
		fmt.Printf("The value %d = 10 x %d\n", x[i], i)
	}
	fmt.Println(len(x))
}