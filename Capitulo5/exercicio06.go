package main

import "fmt"

const (
	_  = iota + 2024
	b
	c
	d
	e
)

func main () {
	fmt.Println(b,c,d,e)
}