package main

import "fmt"

func main() {

	//declara a variavel n do tipo inteiro
	var n int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	//inicializa a variavel soma com 0
	soma := 0
	//soma todos os números de 1 até n
	for i := 1; i <= n; i++ {
		soma += i
	}
	fmt.Printf("A soma dos números de 1 a %d é: %d\n", n, soma)

	//Ex: Digite um número: 5
	//A soma dos números de 1 a 5 é: 15
}
