package main

import "fmt"

func main() {
	t := somentePares(soma, []int{50,51,52,53,54,55,56,57,58,59,60}...)
	fmt.Println(t)
}
 
func soma(x...int) int {
	n := 0 
	for _, value := range x {
		n  += value
	}

	return n
}

func somentePares(f func(x...int) int, y...int) int {
	var slice []int
	for _, value := range y {
		if (value %2 == 0) {
			slice = append(slice, value)
		}
	}
	total := f(slice...)
	return total
}