package main

import "fmt"

// Função f1 sem parâmetros e sem retorno
// Ela apenas imprime "F1" na tela quando chamada
func f1() {
	fmt.Println("F1")
}

// Função f2 que recebe dois parâmetros do tipo string
// Ela imprime os dois parâmetros com uma formatação específica
func f2(p1 string, p2 string) {
	// A formatação está com um erro: o correto seria "/n" ser "\n" para a quebra de linha
	fmt.Printf("F2: %s %s\n", p1, p2)
}

// Função f3 sem parâmetros que retorna uma string
// Ela retorna o valor "F3"
func f3() string {
	return "F3"
}

// Função f4 que recebe dois parâmetros do tipo string e retorna uma string
// Ela retorna uma string formatada usando o fmt.Sprintf
func f4(p1, p2 string) string {
	// Retorna uma string formatada com os valores de p1 e p2
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

// Função f5 sem parâmetros que retorna duas strings
// Ela retorna dois valores de tipo string
func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	// Chama a função f1 que apenas imprime "F1"
	f1()

	// Chama a função f2 passando "Param1" e "Param2" como parâmetros
	// A função imprime "F2: Param1 Param2/n" (lembrando que o '/n' está incorreto)
	f2("Param1", "Param2")

	// Chama a função f3 e f4
	// A função f3 retorna "F3" e f4 retorna uma string formatada com "Param1" e "Param2"
	// A atribuição usa múltiplas variáveis para armazenar os resultados
	r3, r4 := f3(), f4("Param1", "Param2")

	// Imprime o retorno de f3 e f4
	fmt.Println(r3) // Imprime "F3"
	fmt.Println(r4) // Imprime "F4: Param1 Param2"

	// Chama a função f5 e recebe os dois valores retornados
	// Armazena os dois valores retornados em r51 e r52
	r51, r52 := f5()

	// Imprime os valores retornados de f5
	fmt.Println("F5:", r51, r52) // Imprime "F5: Retorno 1 Retorno 2"
}
