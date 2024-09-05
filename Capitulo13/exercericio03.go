package main

import "fmt"

func main() {
	defer vemDepois()
	fmt.Println("Vem  depois")
}


func vemDepois () {
	fmt.Println("Este comentario vem agora")	
}