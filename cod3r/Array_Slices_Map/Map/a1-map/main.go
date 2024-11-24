package main

import "fmt"

func main() {

	/*

	Map, estrutura de dados usada para armazenar pares de chave-valor, onde cada chave única está associada a um valor.
	O map é utilizado quando você precisa de uma estrutura de dados que permite buscar valores rapidamente com base em chaves específicas.
	Ele é eficiente para armazenar e acessar dados quando a chave de busca é conhecida.

					Declara um map sem inicializar.
					var m map[string]int

					inicialização:
					m = make(map[string]int)
					Declaração e inicialização

					m := map[string]int{
				    "Alice": 25,
				    "Bob":   30,
				}
	*/

	aprovados := make(map[int]string)

	aprovados[1234567890] = "Maria"
	aprovados[4726837462] = "João"
	aprovados[9842983429] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

}
