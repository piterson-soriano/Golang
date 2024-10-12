package main

import "fmt"

func main() {
	/*
		x := "oi bom dia. \n Tudo bem? \n espero que sim"
		fmt.Printf(x)
	*/

	x := "Olá!"
	y := " Tudo bem?"

	z := fmt.Sprint(x, y)
	fmt.Println(z)

}
