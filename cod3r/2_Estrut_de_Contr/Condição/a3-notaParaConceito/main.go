package main

import "fmt"

/*
A função notaParaConceito recebe um número(parâmetro n) do tipo float64 e retorna uma string
que retorna o conceito correspondente a essa nota.
*/
func notaParaConceito(n float64) string {

	//Para esse estilo de condição, recomendável utilizar switch para multipla seleções
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
