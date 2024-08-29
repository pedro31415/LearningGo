package main

import "fmt"

func main() {
	fmt.Println("Primeiro Exercicio: ")
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("Este é o valor de x -> ", x)
	fmt.Println("Este é o valor de y -> ", y)
	fmt.Println("Este é o valor de z -> ", z)
	fmt.Printf("%s é uma pessoa bacana que possui %d anos e é uma pessoa %t", y,x,z)	
}