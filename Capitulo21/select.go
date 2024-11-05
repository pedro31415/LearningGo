package main

import "fmt"

func main(){

	fmt.Println("Hello World")

	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <-i
		}
	} (x/2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <-i
		}
	} (x/2)


	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("The value A is ", v)
		case v := <-b:
			fmt.Println("The value B is ", v)
		}
	}
}