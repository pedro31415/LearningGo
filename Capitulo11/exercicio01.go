package main

import "fmt"


type person struct {
	name string
	lastName string
	iceCreamFlavors []string
}

func main() {
	pessoa := person {
		name: "Pedro Henrique",
		lastName: "Araujo Cardoso",
		iceCreamFlavors: []string {
			"flocos",
			"chocolate",
			"mentas",
		},
	}

	pessoa2 := person {
		name: "Carlos Alessandro",
		lastName: "Rubens Dias",
		iceCreamFlavors: []string {
			"flocos-neve",
			"chocolate-branco",
			"sucrilhos",
		},
	}

	fmt.Println("My name is ",  pessoa.name, " ", pessoa.lastName)
	for _,v := range pessoa.iceCreamFlavors {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println("My name is ",  pessoa2.name, " ", pessoa2.lastName)
	for _,v := range pessoa2.iceCreamFlavors {
		fmt.Println(v)
	}
}