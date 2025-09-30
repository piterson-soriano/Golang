package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao() //chama a função
	for {
		exibeMenu() //chama a função

		/*nome, _ := devolveNomeEIdade() //Chama a função e atribui o retorno a variável nome
		fmt.Println(nome)*/

		comando := lerComando() //ler comando do usuário

		switch comando {
		case 1:

			iniciarMonitoramento()
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
}

/*
	func devolveNomeEIdade() (string, int) {
		nome := "Pit"
		idade := 30
		return nome, idade
	}
*/
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

func iniciarMonitoramento() {
	fmt.Println("Monitoramento iniciando...")
	site := "https://httpbin.org/status/404"
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}
