package main

import "fmt"


func main() {
	person := map[string][]string {
		"Cardoso_Pedro": [] string {
			"Leitura",
			"Academia",
			"Estudar",
		},
		"Brilhante_Diana": [] string {
			"Rodar", 
			"Descansar",
			"Tomar banho",
		},
	}

	person["Alessandro_Carlos"] = []string {
		"JS",
		"Trabalhar",
		"Estudar",
	}

	delete(person, "Brilhante_Diana")

	for i,v := range person{
		fmt.Println(i)
		for i, h := range v {
			fmt.Println("\t", i, " -> ", h)
		}
	}
}