package main

import "fmt"

func main () {
	var age int = 21

	if (age >= 18) {
		fmt.Println("You are adult")
	}  else if(age < 0) {
		fmt.Println("The number is invalid age")
	} else {
		fmt.Println("You are young")
	}
}