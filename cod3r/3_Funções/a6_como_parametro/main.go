package main

import "fmt"

// Função que recebe dois inteiros e retorna o produto entre eles
func multiplicacao(a, b int) int {
	return a * b
}

// Função que recebe outra função (que recebe dois inteiros e retorna um inteiro)
// e dois inteiros como parâmetros. Ela executa a função passada com esses dois inteiros.
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	// Aqui, a função multiplicacao é passada como argumento para exec, junto com os inteiros 3 e 4.

	resultado := exec(multiplicacao, 3, 4) // O resultado será 3 * 4 = 12

	fmt.Println(resultado)
}
