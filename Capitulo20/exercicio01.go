package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	
	go firstRoutine()
	go secondRoutine()

	wg.Wait()
}


func firstRoutine(){
	fmt.Printf("Está é uma gourotine")
	wg.Done()
}

func secondRoutine() {
	fmt.Printf("Está é a segunda gourotine")
	wg.Done()
}