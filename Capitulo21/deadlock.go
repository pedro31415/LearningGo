package main

import "fmt"

func main() {
	canal := make(chan int)

	go myLoop(10, canal)

	for v := range canal {
		fmt.Println("Recebido do canal: ", v)
	}

}


func myLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
}