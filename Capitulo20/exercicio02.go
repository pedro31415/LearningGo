package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type humano interface {
	falar() 
}

func (a *pessoa) falar() {
	fmt.Println("Ola meu nome é ", a.nome)
}

func dizerAlgumaCoisa(pessoa humano) {
	pessoa.falar()
}

func main() {
	pessoa1 := pessoa{
		nome: "pedro",
		idade: 21,
	}

	dizerAlgumaCoisa(&pessoa1)
}