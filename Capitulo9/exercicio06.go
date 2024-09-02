package main 


import "fmt"

func main () {
	estados := make([]string, 4, 26)
	estados[0] = "Maranhão"
	estados[1] = "São Paulo"
	estados[2] = "Rio de Janeiro"
	estados[3] = "Amazonas"
	fmt.Println(estados, len(estados), cap(estados))
}