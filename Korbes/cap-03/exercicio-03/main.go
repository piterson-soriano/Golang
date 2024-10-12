package main

import "fmt"

/*
Utilizando a solução do exercício anterior:
EM package-level scope, atribua os seguintes valores às variáveis:
	1. Para "x" atribua 42
	2. Para "y" atribua "James Bond"
	3. Para "z" atribua true

Na função main:
	1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
	Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador de declaração.
	2. Demonstrea variável "s".


*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("Nome: %d, Idade: %s, %t", x, y, z)
	fmt.Println(s)

}
