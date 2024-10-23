package main // Define o pacote principal do programa

import "fmt" // Importa o pacote fmt para operações de formatação e impressão

// Função somar recebe dois inteiros (a e b) e retorna a soma deles
func somar(a int, b int) int {
	// Retorna a soma de a e b
	return a + b
}

// Função imprimir recebe um inteiro e imprime seu valor no console
func imprimir(valor int) {
	// Utiliza a função Println do pacote fmt para exibir o valor.
	fmt.Println(valor)
}
