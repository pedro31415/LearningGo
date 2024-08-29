package main

import "fmt"


func main () {
	x := 0
	for x <= 20 {
		x++
		if(x %2 != 0) {
			continue
		}
		if(x == 20) {
			break
		}
		fmt.Println(x)
	}
}