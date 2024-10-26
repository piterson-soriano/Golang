package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// rand - pacote na biblioteca padrão do Go que fornece funções para geração de números aleatórios
	//NewSource é uma função desse pacote que cria uma nova fonte de números aleatórios.
	s := rand.NewSource(time.Now().UnixNano()) //declara e inicializa uma nova chamada "s".Essa linha cria uma nova fonte de números aleatórios (s).
	r := rand.New(s)                           //Aqui, você cria um novo gerador de números aleatórios (r) que usa a fonte que você acabou de criar. Com r, você pode gerar números aleatórios que são baseados na semente definida por s
	return r.Intn(10)                          // Gera um número aleatório entre 0 e n-1
}

func main() {
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}

	fmt.Println(numeroAleatorio()) // Exemplo de uso
}
