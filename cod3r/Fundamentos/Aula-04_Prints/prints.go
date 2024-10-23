package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é ", x)
	// Na variável xs, é atribuído o valor de x, que é convertido em string pela função Sprint.
	//É útil quando você quer trabalhar com dados como se fossem palavras
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	/* %.2f
	%  - indica que está começando um especificador de formato.
	.2 - mostrará duas casas decimais.
	f  - indica o valor a ser formatado é um número de ponto flutuante(float)*/
	fmt.Printf("O valor de x é %.2f.", x)

}
