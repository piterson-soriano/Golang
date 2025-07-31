package main

import "fmt"

// Declara a variável soma e atribui a ela uma função anônima que recebe dois inteiros (a e b)

var soma = func(a, b int) int {
	return a + b //retorna a soma deles
}

func main() {
	// Chama a função soma com os valores 2 e 3, e imprime o resultado: 5
	fmt.Println(soma(2, 3))

	// Declara uma nova função anônima chamada sub, que faz a subtração de dois inteiros
	sub := func(a, b int) int {
		return a - b
	}
	// Chama a função sub com os valores 2 e 3, e imprime o resultado: -1
	fmt.Println(sub(2, 3))

}
