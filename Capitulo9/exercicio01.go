package main


import "fmt"


func main () {
	value := [5]int {1,2,3,4,5}
	for i, v := range value {
		fmt.Printf("the index -> %d | the value -> %d\n", i, v)
	}
	fmt.Printf("Type: %T", value)
}