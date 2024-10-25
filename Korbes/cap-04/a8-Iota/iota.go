package main

/*
Essa instrução explica o uso da iota em declarações constantes no Go.
A iota é um identificador predefinido que representa constantes inteiras não tipadas sucessivas.

*/
import "fmt"

const (
	// retorna valores sucessivos inteiros não tipados

	a  = iota      // a == 1
	b  = iota      // b == 2
	_  = iota      // ignora o valor 1, não atribuindo a nenhuma constante utilizável
	c0 = iota      // c0 == 3
	_  = iota      // não retorna
	c2 = iota      // c1 == 5
	x  = iota * 10 // x == 60
	y  = iota      // x == 7
	z  = iota * 10 // x == 80
)

func main() {

	fmt.Println(a, b, c0, c2, x, y, z)
}
