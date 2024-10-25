package main

import "fmt"

func imprimirResultado(nota float64) {
	//estrutura básica de condição
	if nota >= 7 { //deve retornar verdadeiro ou falso
		fmt.Println("Aprovado com nota: ", nota)
	} else {
		fmt.Println("Reprovado com nota: ", nota)
	}
}

func main() {
	//chamando a função imprimirResultado
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
