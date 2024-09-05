package main

import "fmt"


type person1 struct {
	name string
	lastName string
	iceCreamFlavors []string
}

func main() {

	// var meumap map[string]person1
	meumap := make(map[string]person1)

	meumap["Araujo Cardoso"] = person1{
		name: "Pedro Henrique",
		lastName: "Araujo Cardoso",
		iceCreamFlavors: []string {
			"flocos",
			"chocolate",
			"mentas",
		},
	}

	meumap["Rubens Dias"] = person1{
		name: "Carlos Alessandro",
		lastName: "Rubens Dias",
		iceCreamFlavors: []string {
			"flocos-neve",
			"chocolate-branco",
			"sucrilhos",
		},
	}

	for _, value := range meumap {
		fmt.Println(value)
	}
}