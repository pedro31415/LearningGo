package main

import "fmt"

func main() {
	x := 10

	finalSquare := square(10)
	func (x int)  {
		fmt.Println(x * 50)
	} (x)

	y := func (x int) int {
		return (x * 100)
	}

	multiplicadorPorDois := multiplicador(2)
	fmt.Println(multiplicadorPorDois(2))

	println("The value is ", y(x))
	println("The are square: ", finalSquare)
}



func square (x int) int {
	return x*x
}

func multiplicador(fator int) func(int) int {
	return func (valor int) int{
		return valor * fator
	}
}