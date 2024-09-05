package main


import "fmt"


func main() {
	fmt.Println(returnInt())
	fmt.Println(returnIntString())
}


func returnInt() int {
	x := 20
	return x 
}

func returnIntString() (int, string) {
	x := 21
	name:= "Pedro"
	return  x,name
}