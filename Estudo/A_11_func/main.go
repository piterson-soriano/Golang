package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()       //chama a função
	exibeMenu()             //chama a função
	comando := lerComando() //ler comando do usuário

	switch comando {
	case 1:
		fmt.Println("Monitoramento iniciando...")
	case 2:
		fmt.Println("Exibindo Logs...")

	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {

	nome := "Pit"
	versao := 1.1
	fmt.Println("Nome: ", nome)
	fmt.Println("Este programa está na versão ", versao)
	fmt.Println()
}
func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento. ")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // fmt.Scanf 	("%d", &comando)
	fmt.Println("O comando escolhido foi ", comandoLido)

	return comandoLido
}
