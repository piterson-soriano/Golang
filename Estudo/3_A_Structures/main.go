package main

import "fmt"

/*
	Declarano uma Variável. Define variavel name do tipo string
		var name string

	Outro exemplo de declaração de variável
		name, idade := "Pit", 39

*/

func main() {
	name, salario := "Pit", 100
	setName(name)

	// Corrigido: usa a função que retorna dois valores
	newSalary, bonus := addSalaryP(salario, 10)
	fmt.Println("Novo salário: ", newSalary)
	fmt.Println("Bonus: ", bonus)
}

func setName(name string) {
	fmt.Println("Nome: ", name)
}

// int int é retornado dois valores compostos
func addSalaryP(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus

}
