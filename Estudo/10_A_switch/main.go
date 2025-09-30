package main

import "fmt"

func main() {

	nome := "Pit"
	versao := 1.1

	fmt.Println("Nome: ", nome)
	fmt.Println("Este programa está na versâo ", versao)

	fmt.Println("1 - Iniciar monitoramento. ")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")

	var comando int
	fmt.Scan(&comando) // fmt.Scanf 	("%d", &comando)

	fmt.Println("O comando escolhido foi ", comando)

	switch comando {
	case 1:
		fmt.Println("Monitoramento iniciando...")
	case 2:
		fmt.Println("Exibindo Logs...")

	case 3:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}
}
