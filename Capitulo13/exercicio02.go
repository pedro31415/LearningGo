package main

import "fmt"

func main() {
	valores := []int {10,20,30,40,50,51,52,53,54,55}

	x := somatorio(valores...)
	y := somatorio2(valores)
	fmt.Println(x, y) 
}

func somatorio(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

func somatorio2(x []int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}