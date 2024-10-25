package main

import "fmt"

func main() {
	numero := 10

	// Verifica se o valor da variável 'numero' é divisível por 2
	if numero%2 == 0 {
		fmt.Println("O número é par.") //se for divisível será par
	} else {
		fmt.Println("O número é impar") //se não, impar.
	}
}
