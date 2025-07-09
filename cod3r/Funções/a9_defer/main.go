package main

/*
O defer em Go é usado para adiaa a execução de uma função até o final da função atual.
Ou seja, ele guarda o que você mandou fazer e só executa quando a função termina.

Onde usar:

Fechamento de arquivos (file.Close())
Liberação de recursos (ex: unlock, disconnect)
Log de finalização
Cleanup (limpeza) em funções que retornam antecipadamente por erro
*/
import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor muito alto!")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))

}
