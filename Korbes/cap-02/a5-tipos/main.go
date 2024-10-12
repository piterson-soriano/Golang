package main

import "fmt"

/*

Tipos em Go são extremamente importantes. (Veremos mais quando chegarmos em métodos e interfaces.)
Tipos em Go são estáticos.
Ao declarar uma variável para conter valores de um certo tipo, essa variável só poderá conter valores desse tipo.
O tipo pode ser deduzido pelo compilador:
x := 10
var y = "a tia do batima"
Ou pode ser declarado especificamente:
var w string = "isso é uma string"
var z int = 15
na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
int, string, bool
Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
slice, array, struct, map
O ato de definir, criar, estruturar tipos compostos chama-se composição. Veremos muito disso futuramente.

*/

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}
