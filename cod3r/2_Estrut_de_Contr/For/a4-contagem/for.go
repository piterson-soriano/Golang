package main

import "fmt"

func main() {
	// Inicializa a variável i com o valor 1
	i := 1

	// Loop que continua enquanto i for menor ou igual a 10
	for i <= 10 {
		// Imprime o valor atual de i
		fmt.Println(i)
		// Incrementa i em 1 a cada iteração
		i++
	}

	// Novo loop com uma nova variável i, iniciando em 0
	for i := 0; i <= 10; i += 1 {
		// Imprime o valor atual de i, com uma nova linha
		fmt.Printf("\n%d", i)
	}
}
