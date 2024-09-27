package main

import "fmt"

type pessoa struct {
	nome string
}

func mudeMe(x * string) {
	*x = "Antonio"
}

func main() {
	pessoa1 := pessoa {"Ronaib"}
	fmt.Printf(pessoa1.nome)
	fmt.Println()
	mudeMe(&pessoa1.nome)
	fmt.Printf(pessoa1.nome)
}
