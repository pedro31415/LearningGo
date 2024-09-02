package main

import "fmt"

func main () {
	i := 2003
	num := 0

	for i <= 2024 {
		fmt.Println("Years -> ", i)
		if(i == 2024) {
			fmt.Printf("\nMy age is %d and year is %d\n", num, i)
		}
		i++
		num++
	}
}