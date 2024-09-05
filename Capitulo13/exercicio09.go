package main

import "fmt"


func valu(x int, callback func(int)) {
	callback(x)
}

func main(){
	valor := func (x int)  {
		fmt.Println("The value is ", x)
	}
	valu(7, valor)
}