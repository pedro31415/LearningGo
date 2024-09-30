package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {
	contador := 0
	totaldegoroutines := 10

	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func ()  {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()	
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}