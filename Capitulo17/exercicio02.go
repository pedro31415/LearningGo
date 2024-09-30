package main

import (
	"encoding/json"
	"fmt"
)

type User []struct {
	First string `json:"first"`
	Age int `json:"age"`
}

func main() {
	u := User{}
	sb := `[{"first":"James","age":32},{"first":"Moneypenny","age":27},{"first":"Monalisa","age":34}]`
	
	// valueJson := []byte(sb)

	err := json.Unmarshal([]byte(sb), &u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}