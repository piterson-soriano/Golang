package main

import "fmt"

// Define um novo tipo chamado nota, baseado no tipo float64.
type nota float64

//Método associado ao tipo nota que verifica se o valor está entre dois limites
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

//Método que converte uma nota para conceito (A, B, C, D ou E)
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

//Função principal que executa o programa.
func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
