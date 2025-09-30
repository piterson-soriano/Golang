package main

import "fmt"

func main() {
	name := getName()
	idade := 39.0
	fmt.Println("Oi!", name, idade)

	salarioAtual := 100
	bonus := 20
	novoSalario := addSalary(salarioAtual, bonus)
	fmt.Println("Novo salário:", novoSalario)
}

// Função que retorna uma string
func getName() string {
	return "Piterson"
}

// Função que retorna um int (salário + bônus)
func addSalary(valorAtual int, bonus int) int {
	return valorAtual + bonus
}
