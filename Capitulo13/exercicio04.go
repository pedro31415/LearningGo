package main

import "fmt"

type pessoa1 struct {
	name string
	sobreNome string
	idade int
}

func main() {
	pessoa01 := pessoa1 {
		name: "Pedro",
		sobreNome: "Carlos",
		idade: 33,
	}

	pessoa01.nomeCompleto()
}


func (p pessoa1) nomeCompleto () {
	fmt.Println("Ola meu nome é ", p.name, p.sobreNome)
}