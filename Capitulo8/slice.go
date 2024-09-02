package main


import "fmt"

func main () {
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	valorTotal := 0

	// add to new value
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	// Range in slice 
	slice3 := []string{"Pedro", "Carla", "Larissa"}
	for index, value := range slice3 {
		fmt.Printf("No indice: %d\nTemos o valor: %s\n", index, value )
	}
	for _, valor := range slice {
		valorTotal += valor
		fmt.Println(valorTotal)
	}
	fmt.Println(valorTotal)

	// slices in array
	sabores := []string{"Portuguesa", "Calabresa", "FrangoCat", "Peporoni"}
	fatia := sabores[1:2]
	fmt.Println(fatia)
	for i:= 0; i< len(sabores); i++ {
		fmt.Printf("The flavor is %s\n", sabores[i])
	}

	// remove element
	sabores = append(sabores[1:])
	fmt.Println(sabores)

	// usign func append

	realSlice := []int{1,2,3,4,5}
	outraSlice := []int{6,7,8,9}

	fmt.Println(realSlice)
	realSlice = append(realSlice, outraSlice...)
	fmt.Println(realSlice)

	// func make 
	slice4 := make([]int, 3, 5)
	for i:=0; i <3; i++ {
		slice4[i] += i + 1
		// fmt.Println(slice4[i])
	}
	fmt.Printf("\n")
	slice4 = append(slice4, 4, 5, 6)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// matrix in slice
	ss := [][]int {
		[]int{1,2,3},
		[]int{4,5,6}, 
		[]int{7,8,9},
	}

	fmt.Println(ss, len(ss), ss[2][2])

	// func maps -> add values in order
	friend := map[string]int {
		"Pedro": 98989898,
		"Rogerio": 7878787,
	}

	// add new value in friend using sintex of map
	friend["Carlos"] = 93993029
	friend["Zahid"] = 90392043

	delete(friend, "Carlos")
	fmt.Println(friend)

	sera, ok := friend["ghost"]
	fmt.Println(friend)
	fmt.Println(sera, ok)


}
