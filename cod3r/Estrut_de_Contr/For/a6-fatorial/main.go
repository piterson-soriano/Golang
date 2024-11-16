package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	//Inicializa o fatorial com o valor 1
	fatorial := 1

	//calcula o fatorial usando um laço for
	for i := 1; i <= n; i++ {
		fatorial *= i
	}
	fmt.Printf("O fatorial de %d é: %d possibilidades", n, fatorial)
}
