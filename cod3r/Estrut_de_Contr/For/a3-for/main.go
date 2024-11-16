package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	// O primeiro laço imprime os números de 1 a 10.
	// Em Go, não existe um comando "while", mas este for simula um comportamento similar.
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// O segundo laço imprime os números pares de 0 a 20.
	// A sintaxe do for é utilizada de forma tradicional, incrementando de 2 em 2.
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}
	fmt.Println("\n############################")

	// O terceiro laço verifica se os números de 1 a 10 são pares ou ímpares.
	for i := 1; i <= 10; i++ {
		// Se o resto da divisão de i por 2 for igual a 0, imprime "Par"; caso contrário, "Impar".
		if i%2 == 0 {
			fmt.Print("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// O último laço é um laço infinito que imprime "Para sempre..." a cada segundo.
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
