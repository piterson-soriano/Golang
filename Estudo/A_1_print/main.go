package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		os.Args
		ARGS[0,1]
	*/
	// Verifica se o número de argumentos passados na linha de comando é diferente de 2.
	if len(os.Args) != 2 {
		// Se não houver exatamente 1 argumento passado (além do nome do programa),
		// o programa sai com status de erro (código 1).
		os.Exit(1)
	}
	fmt.Println("Teste", os.Args[1])
}
