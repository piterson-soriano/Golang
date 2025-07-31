package main

import "fmt"

func main() {
	// Inicializa a variável i com o valor 1
	i := 1

	// Primeiro laço: Enquanto i for menor ou igual a 10
	for i <= 10 {
		fmt.Println(i) // Imprime o valor atual de i
		i++            // Incrementa i em 1 a cada iteração
	}

	// Segundo laço: Inicializa uma nova variável i (escopo local) com 0
	for i := 0; i <= 10; i += 1 {
		fmt.Printf("\n%d", i) // Imprime o valor de i seguido de uma nova linha
	}
}
