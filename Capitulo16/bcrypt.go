package main

import "fmt"
import "golang.org/x/crypto/bcrypt"

func main() {
	senha := "26março2002"
	senhaErrada := "26março2000"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaErrada)))
}