package main

import "fmt"

func main() {

	/*
		Um ponteiro em Go é uma variável que armazena o endereço de outra variável, em vez de armazenar diretamente um valor
		Um dos principais usos de ponteiros é permitir que funções mudem o valor de variáveis fora de seu escopo.
		Se você passar uma variável normalmente para uma função, ela recebe uma cópia. Com ponteiros, a função pode alterar o valor original.
	*/

	num := 42

	// Crie um ponteiro chamado "valor" que aponta para a variável num
	var valor *int = &num

	fmt.Println("Valor original:", num)
	fmt.Println("Endereço de número:", valor)
	fmt.Println("Valor através do ponteiro:", *valor)

	*valor = 100

	fmt.Println("Valor modificado:", num)

}
