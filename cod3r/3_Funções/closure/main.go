package main

import "fmt"

/*
Closure é uma função que captura o ambiente onde foi criada.
Isso significa que ela pode acessar variáveis do escopo onde foi definida.
Usado para:
Para encapsular estado sem usar structs.
Para criar funções geradoras (como contadores, acumuladores).
Em callbacks ou funções que precisam de contexto.
*/

// closure retorna uma função que acessa uma variável do seu escopo externo
func closure() func() {
	x := 10 // Variável local à função closure

	var funcao = func() {
		// Essa função anônima é uma closure: ela "captura" a variável x acima
		fmt.Println(x) // Vai imprimir o valor de x que está no escopo da função closure (ou seja, 10)
	}
	return funcao // Retorna a função anônima, preservando o acesso à variável x
}

func main() {
	x := 20        // Declara uma variável x local ao escopo da função main
	fmt.Println(x) // Imprime 20

	imprimeX := closure() // Chama a função closure, que retorna a função anônima
	imprimeX()            // Executa a função retornada; imprime 10, que é o valor capturado pela closure
}
