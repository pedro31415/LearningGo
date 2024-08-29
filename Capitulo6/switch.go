package main

import "fmt"


func main () {
	x := 100 - 200

	 switch {
	 	case x == 0:
			fmt.Println("The value is less than ", x)
			break
		case x == 10:
			fmt.Println("The value equal to ", x)
			break
		case x > 10:
			fmt.Println("The value is greater than ", x)
			break
		default:
			fmt.Println("Number invalid")
	 }

	 switch x {
	 	case -100:
			fmt.Println("TRUE VALUE", x)
		case 300: 
			fmt.Println("IT IS NOT TRUE", x)
		default:
			fmt.Println("INVALID NUMBER")
	 }
}