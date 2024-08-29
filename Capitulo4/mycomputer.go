package main

import (
	"fmt"
	"runtime"
)

//Estudo de constantes

const value = 10
const (value2 = 10)

// Iota
const (
	a = iota + 10 * 4
	_
	a2
	a3
	a4
)

func main () {
	name := "Pedro Henrique Araujo Cardoso"
	sliceName := []byte(name)

	fmt.Println(value, value2)
	fmt.Println(a,a2,a3,a4)

	fmt.Printf("The name: %v\n The formated: %T\n", sliceName, sliceName)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(" ")
	s := "ascii éøâ 香"
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

	y := 24
	x := y << 2
	z := y >> 2

	fmt.Printf("The value -> %d = %b\n", y, y)
	fmt.Printf("The value -> %d = %b\n", x, x)
	fmt.Printf("The value -> %d = %b\n", z, z)
}