package main

import "fmt"

/*
Use var para declarar três variáveis. Elas dever package-level scope.
Nâo atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis.

1. identificador "x" deverá ter tipo int
2. identificador "y" deverá ter tipo string
3. identificador "z" deverá ter tipo bool
*/
var x int
var y string
var z bool

func main() {

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
}
