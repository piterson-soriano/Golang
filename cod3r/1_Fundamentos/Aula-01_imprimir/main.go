package main // Declara que este arquivo pertence ao pacote "main", que é o pacote principal onde a execução do programa começa.

// Importa o pacote "fmt", que contém funções para formatar e imprimir texto, entre outras coisas.
import "fmt"

// Define a função principal "main", que é o ponto de entrada do programa.
func main() {
	// Usa a função Print do pacote "fmt" para imprimir "Primeiro" na saída padrão, sem pular linha.
	fmt.Print("Primeiro")
	// Usa a função Print novamente para imprimir "Programa!" na mesma linha, logo após "Primeiro".
	fmt.Print(" programa!")
} // Fecha a definição da função main.
