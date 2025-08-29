package main

import (
	"fmt"
	"os"
)

// Define o tipo personalizado 'nota'
type nota float64

// Método que verifica se a nota está entre dois valores
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

// Função que converte nota em conceito
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

// Função principal
func main() {
	var entrada float64

	for {
		fmt.Print("Digite uma nota entre 0.0 e 10.0 (ou -1 para sair): ")
		_, err := fmt.Scan(&entrada)

		// Verifica se houve erro na leitura
		if err != nil {
			fmt.Println("Erro ao ler a entrada. Tente novamente.")
			continue
		}

		// Condição de saída
		if entrada == -1 {
			fmt.Println("Encerrando o programa.")
			os.Exit(0)
		}

		// Validação da nota
		if entrada < 0.0 || entrada > 10.0 {
			fmt.Println("Nota inválida. Digite um valor entre 0.0 e 10.0.")
			continue
		}

		// Converte a nota e imprime o conceito
		n := nota(entrada)
		conceito := notaParaConceito(n)
		fmt.Printf("Nota %.2f => Conceito %s\n\n", entrada, conceito)
	}
}
