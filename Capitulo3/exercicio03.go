package main


import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main () {
	s := fmt.Sprintf("%s é uma pessoa bacana que possui %d anos e é uma pessoa %t", y,x,z)
	fmt.Println(s)
}