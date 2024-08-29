package main

import "fmt"

func main () {
	x := 140

	if (x <= 150) || (x >= 200) {
		fmt.Println("THe value is true")
	} else {
		fmt.Println("Erro in the server")
	}
}