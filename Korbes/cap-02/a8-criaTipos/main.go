package main

import (
	"fmt"
	"io"
	"os"
)

/*type hotdog int

var v hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v\n", x)

}*/

func main() {
	const nome, idade = "Pit", 22
	s := fmt.Sprintf("%s is %d years old.\n", nome, idade)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
