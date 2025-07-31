package main

import "fmt"

// Função que recebe dois inteiros e retorna dois inteiros com nomes
func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2  // 'segundo' recebe o valor de p2
	primeiro = p1 // 'primeiro' recebe o valor de p1
	return        // Retorna os valores nomeados (segundo, primeiro)
}

func main() {
	x, y := trocar(2, 3) // Chama a função com 2 e 3; x=3, y=2
	fmt.Println(x, y)    // Imprime: 3 2

	segundo, primeiro := trocar(7, 1) // Chama a função com 7 e 1; segundo=1, primeiro=7
	fmt.Println(segundo, primeiro)    // Imprime: 1 7
}
