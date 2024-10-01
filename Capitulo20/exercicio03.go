package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
const quantidadeDeGoroutines = 200
var contador int

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Goroutines: ", quantidadeDeGoroutines, contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}