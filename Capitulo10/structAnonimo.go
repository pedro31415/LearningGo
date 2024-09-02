package main

import "fmt"


func main() {
	x := struct {
		name string
		age int
	} {
		name: "Romario",
		age: 82,
	}

	fmt.Println(x)
}