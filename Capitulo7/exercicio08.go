package main

import "fmt"

func main () {
	myAge := 200
	myName := "Pedro"

	switch myAge {
		case 21:
			fmt.Println("My age 21")
		case 18:
			fmt.Println("It is not my age")
		case 20: 
			fmt.Println("little more ")
		default:
			fmt.Println("The number is  invalid for to  question")
	}

	switch {
		case myName == "Pedro":
			fmt.Println("This is my name Pedro")
		case myName == "Carlos":
			fmt.Println("This is my friend name")
		case myName == "Cristiane":
			fmt.Println("This is my mom name")
		default:
			fmt.Println("Invalid name")
	}
}