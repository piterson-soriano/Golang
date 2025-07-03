package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		// Imprime cada nome precedido de um número de ordem, começando por 1
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "José", "Guilherme"}
	// Chama a função imprimirAprovados, usando o operador "..." para expandir a slice em argumentos individuais
	imprimirAprovados(aprovados...)
}
