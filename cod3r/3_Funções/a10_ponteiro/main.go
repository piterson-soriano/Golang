package main

import "fmt"

/*ponteiros são usados para armazenar o endereço de memória de uma variável,
permitindo acessar e modificar diretamente o valor dessa variável fora do
seu escopo original (por exemplo, dentro de funções).*/

func inc1(n int) {
	n++ // é o mesmo que n = n + 1
}

// revisão: um ponteiro é uma referência para o valor na memória.
// é representado por um asterisco (*) antes do tipo.
func inc2(n *int) {
	// O * é usado para desreferenciar o ponteiro, ou seja, acessar o valor que ele aponta.
	*n++
}

func main() {
	n := 1
	inc1(n)        // passa o valor de n, não o endereço na memória
	fmt.Println(n) //imprime 1, pois inc1 não altera o valor original
}
