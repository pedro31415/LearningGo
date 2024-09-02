package main

import "fmt"

type cliente struct {
	name string
	lastName string
	phoneNumber int
}

func main() {
	cliente1 := cliente {
		name: "Pedro Henrique",
		lastName: "Araujo Cardoso",
		phoneNumber: 9908024929,
	}

	fmt.Println(cliente1)
}