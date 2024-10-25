package main

import "fmt"

/*
	constantes são valores imutáveis
	Podem ser tipadas ou não;
		const oi = "bom dia"
		const oi string = "bom dia"

	As não tipadas só terão um tipo atribuido a elas quando forem usadas.const
		ex: Qual o tipo de 42? int? uint? float64?
		ou seja, é uma flexibilidade conveniente
	Na prática: int, float, string.
		const x = y
		const(x = y)
*/

const x = 10 //tipo da constante só é definido no momento do uso

const (
	a = 10
	b = 20
	c = 30
)

var y = 10 // tipo da variavel, é definido no momento da atribuição

var e int
var d float64

func main() {
	d = x
	fmt.Println(d)
	fmt.Println(a, b, c)
}
