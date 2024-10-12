package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) //Valor Zero
	x = true
	fmt.Println(x) //Valor atribuido
	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)
	fmt.Println("x: 10 < 100 = ", x)
	fmt.Println("y: 10 == 100 = ", y)
	fmt.Println("z: 10 > 100 = ", z)
}
