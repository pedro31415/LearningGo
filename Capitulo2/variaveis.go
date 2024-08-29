package main

import "fmt"

var y float32
var z string
type hotdog int
var b hotdog = 101

func main(){

	fmt.Printf("%v -> %T\n", y, y);
	fmt.Printf("%v -> %T\n", z, z);
	fmt.Printf("%v -> %T\n", b, b);

	x := 10 
	x = int(b)
	fmt.Println(x)
}

