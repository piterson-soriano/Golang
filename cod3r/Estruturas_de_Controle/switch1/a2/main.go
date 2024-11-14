package main

import "fmt"

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("Caso 1")
	case 2:
		fmt.Println("Caso 2")
		fallthrough
	case 3:
		fmt.Println("Caso 3")
	case 4:
		fmt.Println("Caso 4")
	default:
		fmt.Println("Caso default")
	}
}
