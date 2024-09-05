package main

import "fmt"

func main() {
	pessoa := struct {
		name string
		hobbys []string
		trabalhos map[string]int
	} {
		name: "Pedro Henrique Araujo Cardoso",
		hobbys: []string {
			"Correr",
			"jogar futsal",
			"tenis",
		},
		trabalhos: map[string]int{
			"Corretor": 1,
			"Personal": 2,
		},
	}

	fmt.Println(pessoa)
}