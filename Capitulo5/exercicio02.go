package main

import "fmt"

func main () {
	value := 20
	value2 := 40
	z := (value == value2)
	n := (value >= value2)
	m := (value <= value2)
	c := (value > value2)
	p := (value < value2)
	pedro := (value != value2)

	fmt.Println(z,n,m,c,p,pedro)
}