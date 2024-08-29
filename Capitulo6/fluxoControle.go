package main


import "fmt"


func main () {
	//Estrutura Inicial de um for
	// for x:= 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }
	// RELOGIO 
	for hours := 0; hours <= 12; hours++ {
		for minutes:= 0; minutes < 60; minutes++ {
			fmt.Printf("%d:%d\n", hours, minutes)
		}
	}
}