package main

import "fmt"

func main() {
	/*
		Declarano uma Variável. Define variavel name do tipo string
			var name string

		Outro exemplo de declaração de variável
			name, idade := "Pit", 39

	*/

	name, salario := "Pit", 100
	setName(name)
	//Adiciona mais "R$10.00"
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo salário: ", newSalary)
	fmt.Println("Bonus: ", bonus)

}

func setName(name string) {
	fmt.Println("Nome: ", name)
}

func addSalary(p Pessoa, bonus int) {
	//o comando p.salario = p.salario + bonus é o mesmo que o comando da linha abaixo
	p.salario += bonus

}

// int int é retornado dois valores compostos
func addSalaryP(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus

}
