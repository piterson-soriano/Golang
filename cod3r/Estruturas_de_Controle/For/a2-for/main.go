package main

import "fmt"

func main() {
	soma := 0

	// A condição do loop foi corrigida para verificar 'i <= 100'
	for i := 1; i <= 100; i++ {
		soma += i // Soma i à variável soma
	}
	fmt.Println("A soma dos números de 1 a 100 é:", soma)
}
