package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string `json:"first"`
	Age int `json:"age"`
}


func main() {
	u1 := User{
		First: "James",
		Age: 32,
	}

	u2 := User {
		First: "Moneypenny",
		Age: 27,
	}

	u3 :=User {
		First: "Monalisa",
		Age: 34,
	}

	users := []User{u1,u2,u3}
	fmt.Println(users)

	userData, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Println(string(userData))
}