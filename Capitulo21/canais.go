package main

import "fmt"


func main() {
	canal := make(chan int)

	go send(canal)

	receive(canal)

}


func send(s chan int) {
	s <- 42
}

func receive(s chan int) {
	fmt.Println("O valor recebido do canal ", <-s)
}