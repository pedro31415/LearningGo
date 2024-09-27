package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int {19,34,100,1,33,2}
	slice2 := []string{"Arnaldo", "Gauvao bueno", "Casa Grande", "Ronaldo Fenomeno"}

	fmt.Println(slice2)
	sort.Strings(slice2)
	fmt.Println(slice2)
	fmt.Println()
	fmt.Println(slice1)
	sort.Ints(slice1)
	fmt.Println(slice1)
}