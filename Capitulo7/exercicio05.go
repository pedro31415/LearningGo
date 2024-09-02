package main

import "fmt"


func main () {
	for i := 10; i <= 100; i++ {
		num := i % 4 
		fmt.Printf("The rest divisible %d for 4 is %d\n", i, num)
	}
}