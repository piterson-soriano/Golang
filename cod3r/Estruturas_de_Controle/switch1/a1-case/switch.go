package main

import "fmt"

func notaParaConceito(n float64) string {

	var nota = int(n)
	switch nota {

	case 10:
		/*
			fallthrough: Quando desejar que o código continue para o próximo case.
				Se não usar, o comportamento padrão é que a execução saia do switch após o primeiro case correspondente.
		*/
		//fallthrough
		return "X"
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(10.1))
}
