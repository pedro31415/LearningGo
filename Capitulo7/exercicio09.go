package main

import "fmt"

func main () {
	esporteFavorito := "Futebol"

	switch {
	case esporteFavorito == "Handeboll":
		fmt.Println("This is my favorite sport Handeboll")
	case esporteFavorito == "Futebol":
		fmt.Println("This is my favorite sport Futebol")
	case esporteFavorito == "Volei":
		fmt.Println("This is my favorite sport Volei")
	default:
		fmt.Println("this is not favorite sport")
	}
}