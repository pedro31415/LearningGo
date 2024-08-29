package main

import "fmt"


func main () {
	var x int = 144
	fmt.Printf("The value Numerical -> %d\nThe value Binary -> %b\nThe value Hexadecimal -> %#x\n\n", x,x,x)
	
	var y int = x << 1
	fmt.Printf("The value Numerical -> %d\nThe value Binary -> %b\nThe value Hexadecimal -> %#x\n", y,y,y)
}