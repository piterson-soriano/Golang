package main

import "fmt"

/*
Definiu ama função obterResultado. A função recebeu recebe uma nota do tipo float64.
Retorna uma string representando o resultado correspondente a esssa nota.
*/
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println("Nota: 6.2")
	fmt.Print("Resultado: ")
	fmt.Println(obterResultado(6.2))

}
