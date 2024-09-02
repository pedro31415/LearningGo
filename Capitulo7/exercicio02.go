package main 

import "fmt"

func main () {
	for i := 65; i <= 90; i++ {
		for z := 0; z < 3; z++ {
			fmt.Printf("[%d] --> %c\n", i, i)
		}
		fmt.Println()
	}
}