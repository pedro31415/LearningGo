package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	defer fmt.Println("Rodou a função que estava em defer.")

	r := strings.NewReader("Luan Gameplays")

	io.Copy(f,r)
}