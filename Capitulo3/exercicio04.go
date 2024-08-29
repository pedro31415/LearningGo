package main

import (
	"fmt"
)

type pedro int 
var x pedro
var p int   
func main () {
	phrase := (30 > 20)
	fmt.Println("Este éo valor de x -> ", x)
	fmt.Printf("Este é o tipo da variavel: %T\n", x)
	x = 42 
	fmt.Println(x)
	p = int(x)
	fmt.Printf("O valor é igual: %d\nTipo da Variavel: %T\n", p,p)
	fmt.Print(phrase)

}