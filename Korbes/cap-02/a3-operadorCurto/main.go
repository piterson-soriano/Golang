package main

import "fmt"

//package level scope - abrangência do nível ao pacote.
//ou seja a variave y estará disponível em qualquer parte do código

var y = "Teste golang"

/*
Terminologia:
- Keywords -> (palavra chave ) são termos reservados
- Operadores, operandos
- Statement (Declaração, afirmação)
- Expressão -> Qualquer coisa que "produz um resultado"


*/

func main() {

	// := só funciona dentro de code block
	x := 10
	y := "teste"

	// %v -> valor
	// %T -> Tipo de dados

	//float: float in point - Flutuante pq tem ponto
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

}
