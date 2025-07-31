package main

import "fmt"

// Função que calcula a média entre duas notas
func calcularMedia(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}

// Função que verifica se a média é suficiente para aprovação
func verificarAprovacao(media float64) string {
	if media >= 6.0 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	// Declara duas notas
	nota1 := 7.5
	nota2 := 5.0

	// Calcula a média chamando a função
	media := calcularMedia(nota1, nota2)

	// Imprime a média
	fmt.Println("Média:", media)

	// Verifica e imprime o resultado da aprovação
	resultado := verificarAprovacao(media)
	fmt.Println("Resultado:", resultado)
}
