package main

import "fmt"


func main () {
	x := 10
	y := &x

	*y = 20

	incremento(&x)
	fmt.Printf("%p -> x %p -> y\n", &x, &y)
	fmt.Printf("%p -> x %p -> y\n", &x, &*y)
	fmt.Println(*y, "->", x)
}

func incremento(y *int) {
	*y++
}