package main

import "fmt"

func main () {
	valorMultiplicadosporDez := func (x int)  int {
		return x * 10
	}

	result := valorMultiplicadosporDez(2)

	fmt.Println(result)
}