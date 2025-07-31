package main

import "fmt"

func main() {

	// Cria um slice de inteiros com comprimento 10 e capacidade 10.
	// Todos os elementos são inicializados com o valor zero.
	s := make([]int, 10) // "make" aloca memória para o slice.

	// Define o valor 12 na posição de índice 9 do slice.
	s[9] = 12 // Último elemento do slice é atualizado para 12.

	// Imprime o slice completo.
	fmt.Println(s) // Saída: [0 0 0 0 0 0 0 0 0 12]

	// Recria o slice, agora com comprimento 10 e capacidade 20.
	// Comprimento é o número de elementos acessíveis inicialmente (10),
	// e capacidade é o total de espaço alocado (20).
	s = make([]int, 10, 20) // Os primeiros 10 elementos são zeros.

	// Imprime o slice, seu comprimento e capacidade.
	fmt.Println(s, len(s), cap(s)) // Saída: [0 0 0 0 0 0 0 0 0 0] 10 20

	// Adiciona 10 elementos ao slice usando a função "append".
	// Isso aumenta o comprimento do slice, mas não ultrapassa sua capacidade (20).
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	// Imprime o novo slice, seu comprimento e capacidade.
	fmt.Println(s, len(s), cap(s)) // Saída: [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0] 20 20

	// Adiciona mais um elemento ao slice.
	// Agora o comprimento do slice excede sua capacidade atual (20).
	// O runtime aloca um novo array subjacente com capacidade maior (geralmente o dobro da anterior).
	s = append(s, 1)

	// Imprime o slice atualizado, seu comprimento e sua nova capacidade.
	fmt.Println(s, len(s), cap(s)) // Saída: [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 1] 21 40
}
