package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome string
	potencia float64
	consumo float64
}

type ordenarPorPotencia [] carro
type ordenarPorConsumo [] carro
type ordenarPorLucroProDono [] carro

func (x ordenarPorPotencia) Len() int {
	return len(x)
}

func(x ordenarPorPotencia)  Less(i, j int) bool { 
	return x[i].potencia < x[j].potencia
}

func (a ordenarPorPotencia) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (x ordenarPorConsumo) Len() int {
	return len(x)
}

func(x ordenarPorConsumo)  Less(i, j int) bool { 
	return x[i].consumo < x[j].consumo  
}

func (a ordenarPorConsumo) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (x ordenarPorLucroProDono) Len() int {
	return len(x)
}

func(x ordenarPorLucroProDono)  Less(i, j int) bool { 
	return x[i].consumo > x[j].consumo
}

func (a ordenarPorLucroProDono) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	carros := []carro { carro{"chevete", 50, 5}, carro{"porsche", 300, 5}, carro{"fusca", 20, 30}}
	fmt.Println(carros)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)
	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)
	sort.Sort(ordenarPorLucroProDono(carros))
	fmt.Println(carros)
}