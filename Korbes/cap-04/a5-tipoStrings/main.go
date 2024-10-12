package main

import "fmt"

func main() {

	/*	s :=
		`
		..
		...
		....
		.....
		......
		.......
		`
			fmt.Printf("%v\n%T", s, s)
	*/
	s := "Olá, playground"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v)
	}

}
