package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	fmt.Println("-----------------Inteiros-----------------")

	//números inteiros
	fmt.Println("numeros inteiros:", 1, 2, 1000)

	/*
		O comando abaixo imprime a string "Literal inteiro é" seguida do tipo do valor 32000.
		A função reflect.Typeof (do pacote reflect) é utilizada para determinar o tipo do valor passado, que neste caso é um inteiro.
		Assim, o resultado será algo como int ou int32, dependendo do sistema e da arquitetura.


		O comando reflect.TypeOf(b) em Go é usado para obter o tipo de uma variável em tempo de execução.
		A função reflect.TypeOf faz parte do pacote reflect, que fornece funções para inspeção de tipos.
	*/
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	/*
		uint: Este é um tipo que tem tamanho dependente da arquitetura do sistema

		uint8 -  Representa um inteiro sem sinal de 8 bits.  Pode armazenar valores de 0 a 255
		uint16 - Representa um inteiro sem sinal de 16 bits. Pode armazenar valores de 0 a 65.535
		uint32 - Representa um inteiro sem sinal de 32 bits. Pode armazenar valores de 0 a  4.294.967.295
		uint64 - Representa um inteiro sem sinal de 64 bits. Pode armazenar valores de 0 a 18.446.744.073.709.551.615.
	*/

	// Declara uma variável b do tipo byte e atribui o valor 255
	var b byte = 255
	// Imprime o tipo da variável b
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal
	valorMax := math.MaxInt64
	// Imprime o valor máximo de int
	fmt.Println("O valor máximo de int é", valorMax)
	// Imprime o tipo da variável valorMax
	fmt.Println("O tipo de valor máximo é", reflect.TypeOf(valorMax))

	// Declara uma variável i2 do tipo rune e atribui o valor 'a'
	var i2 rune = 'a'
	// Imprime o tipo da variável i2
	fmt.Println("O rune é", reflect.TypeOf(i2))
	// Imprime o valor da variável i2
	fmt.Println("rune", i2)

	// Declara uma variável x do tipo float32 e atribui o valor 49.99
	var x float32 = 49.99
	// Imprime o tipo da variável x
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	// Imprime o tipo do literal 49.99 (que é float64 por padrão)
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	fmt.Println("-----------------boolean-----------------")

	bo := true
	fmt.Print("O tipo de bo é: ", reflect.TypeOf(bo), " -> ")
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Pit"
	fmt.Println(s1, "!")
	fmt.Println("O tamanho da string é: ", len(s1))

	//string com multiplas linhas
	s2 := `Olá 
	meu 
	nome
	é
	Pit`
	fmt.Println("O tamanho da string é", len(s2))

	//char

	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))

}
