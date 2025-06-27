package main

import "fmt"

// Função que calcula a média de uma lista de números float64
func media(numeros ...float64) float64 {
	// Inicializa uma variável para armazenar o total da soma dos números
	total := 0.0
	for _, num := range numeros { // Loop pelos números fornecidos como argumentos
		// Soma cada número ao total
		total += num
	}

	// Retorna a média: total dividido pela quantidade de números
	return total / float64(len(numeros))
}

func main() {
	// Chama a função media com alguns valores e imprime o resultado formatado com 2 casas decimais
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
