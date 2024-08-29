package main


import "fmt" 


func main () {
	rawString := `Essa frase 
	\n não é minha\\`
	fmt.Println(rawString)
}