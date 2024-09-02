package main 

import "fmt"

func main () {
	i := 2003
	for {
		if(i <= 2024) {
			fmt.Println("Years -> ", i)
			i++
		} else {
			break
		}
	}
}