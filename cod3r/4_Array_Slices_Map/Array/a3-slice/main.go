package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declara um array fixo de tamanho 3 com os elementos {1, 2, 3}.
	a1 := [3]int{1, 2, 3} // Array é uma coleção de elementos de tamanho fixo.

	// Declara um slice, que é semelhante a um array, mas com tamanho dinâmico.
	s1 := []int{1, 2, 3} // Slice é uma "janela" flexível sobre arrays.

	// Imprime o conteúdo do array e do slice.
	fmt.Println("a1: ", a1, "s1: ", s1)

	// Usa a biblioteca "reflect" para mostrar o tipo das variáveis.
	// "a1" é do tipo "[3]int" (array de tamanho fixo com 3 inteiros).
	// "s1" é do tipo "[]int" (slice de inteiros).
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Declara outro array fixo com 5 elementos.
	a2 := [5]int{1, 2, 3, 4, 5}

	// Cria um slice "s2" que contém elementos da posição 1 até a posição 2 (não inclusiva).
	s2 := a2[1:2] // Slice pega apenas o elemento da posição 1.
	fmt.Print("s2 - 1:2 ")
	fmt.Println(a2, s2) // Mostra o array original e o slice gerado.

	// Cria um slice "s3" que contém elementos da posição 1 até a posição 3 (não inclusiva).
	s3 := a2[1:3] // Slice pega os elementos da posição 1 e 2.
	fmt.Print("s3 - 1:3 ")
	fmt.Println(a2, s3)

	// Cria um slice "s4" que contém elementos da posição 1 até a posição 4 (não inclusiva).
	s4 := a2[1:4] // Slice pega os elementos da posição 1, 2 e 3.
	fmt.Print("s4 - 1:4 ")
	fmt.Println(a2, s4)

	// Cria um slice "s5" que contém elementos da posição 1 até a posição 5 (não inclusiva).
	s5 := a2[1:5] // Slice pega os elementos da posição 1, 2, 3 e 4.
	fmt.Print("s5 - 1:5 ")
	fmt.Println(a2, s5)

	// Cria um slice "s6" a partir de "s3", mas limitado apenas ao primeiro elemento.
	// "s3[:1]" significa: pegue o primeiro elemento de "s3".
	s6 := s3[:1]
	fmt.Print("s6 - ")
	fmt.Println(s3, s6) // Mostra que "s6" é um subconjunto de "s3".

	// Nota: O slice é apenas uma referência ao array subjacente ("a2" nesse caso).
	// Alterações feitas no slice podem impactar o array original.
}
