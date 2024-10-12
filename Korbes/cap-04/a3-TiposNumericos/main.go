package main

import "fmt"

func main() {

	a := "a"
	b := "á"
	c := "@"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	//conversão do tipo
	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v", d, e, f)
}
