package main

import "fmt"


func main() {
	// wrong
	person := [][]string {
		[]string {
			"Pedro Henrique", "Maria Cristiane", "Diana Brilhante",
		},
		[]string {
			"Araujo Cardoso", "Carvalho Araujo", "Araujo Cardoso",
		},
		[]string {
			"Leitura", "Assistir TV", "Rodar",
		},
	}
	fmt.Println(person)

	for i,v:= range person {
		// for j := 0; j < len(person); i++ {
		// 	fmt.Printf(person[i][j])
		// }
		// fmt.Println(v)
		fmt.Println(i, v)
	}

	
}