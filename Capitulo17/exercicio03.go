package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First string `json: "first"`
	Age int `json: "age"`
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

	users := []User {u1,u2,u3}
	valueJson := json.NewEncoder(os.Stdout)
	err := valueJson.Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}