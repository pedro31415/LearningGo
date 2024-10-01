package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
const quantidadeDeGoroutines = 200
var contador int
var mu sync.Mutex

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Goroutines: ", quantidadeDeGoroutines, contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}