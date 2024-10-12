package main

import "fmt"

/*
Utilizando a solução do exercicio anterior:
  1. Em package-level scope, utilizando a palavra chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo
  subjacente do tipo que você criou no exerc. anterior.
  2. Na função main:
	1. Isto já deve estar feito:
		1. Demonstre o valor da variável "x"
		2. Demonstre o tipo da variável  "x"
		3. Atribua 42 à variável "x" utilizando o operador "="
		4. Demonstre o valor da variável "x"
*/

type valor int

var x valor
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("x: %T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T", y)

}
