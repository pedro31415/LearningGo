package main

import "fmt"


func main() {
	slice := []int {1,2,3,4,5,6,7,8,9,10}
	slice1 := slice[0:3]
	slice2 := slice[4:]
	slice3 := slice[1:7]
	slice4 := slice[2:len(slice) - 1]
	fmt.Print(slice1, slice2, slice3, slice4)

}


// do primeiro a o terceiro 
// do quinto ao ultimo 
// do segundo ao setimo 
// obter o mesmo resultado usando a func len() para determinar o penultimo termo 