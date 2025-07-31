package main

import "fmt"

// Função principal, ponto de entrada do programa.
func main() {
	// Criação de um array de inteiros usando o operador "..." para inferir o tamanho do array.
	numeros := [...]int{1, 2, 3, 4, 5}

	// Primeira iteração sobre o array usando o "range" (que retorna o índice e o valor).
	for i, numero := range numeros {
		// Exibe o índice (i+1 para começar de 1) e o número correspondente do array.
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	// Segunda iteração sobre o array, mas apenas o valor de cada elemento é utilizado (ignorando o índice).
	for _, num := range numeros {
		// Exibe cada número do array em uma nova linha.
		fmt.Println(num)
	}
}
