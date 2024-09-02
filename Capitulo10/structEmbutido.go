package main


import "fmt"


type person struct {
	name string
	age int
}

type profession struct {
	person
	title string
	wage float64
}

func main() {
	pessoa1 := person {
		name: "Pedro",
		age: 21,
	}

	pessoa2 := profession {
		person: person{
			name: "Carlos",
			age: 21,
		},
		title: "Programming master",
		wage: 30000,
	}

	pessoa3 = profession{person{"Ronaldo",54,}, "Socccer player", 20000000000}


	pessoa1.name = "Rogerio do Gelo"

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.name)
}