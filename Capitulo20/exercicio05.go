package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
const quantidadeDeGoroutines = 200
var contador int32 

func main() {
	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Goroutines: ", quantidadeDeGoroutines, "Contador:", contador)
}

func criarGoroutines(i int) {
	wg.Add(i) 
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&contador, 1) 
			runtime.Gosched()             
			v := atomic.LoadInt32(&contador) 
			fmt.Println(v)               
			wg.Done()                    
		}()
	}
}
