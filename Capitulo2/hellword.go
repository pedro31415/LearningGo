package main

import "fmt"

func main() {
	numeroBytess, _ := fmt.Println("Hello Wolrd")
	numeroBytes, erro := fmt.Println("My name is Pedro, and you?!")
	fmt.Println(numeroBytes, erro)
	fmt.Println(numeroBytess)

	y := 10
	qualquerCoisa(y)
}


func qualquerCoisa (x int) {
	fmt.Println(x)
}