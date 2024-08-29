package main 

import "fmt"

func main () {
	x := 90

	if x > 100 {
		fmt.Print("Hello Wolrd")
	} else if x == 90 {
		fmt.Print("The number is equal ", x)
	} else {
		fmt.Print("The number is greater thaan ", x)
	}
}