package main // Define o pacote principal do programa.

import "fmt" // Importa o pacote fmt, usado para entrada e saída formatada.

func main() {
	/*
	   Declarando uma variável:
	   Define a variável name do tipo string:
	       var name string

	   Outro exemplo de declaração de variável com inferência de tipo:
	       name, idade := "Pit", 39
	*/

	// Declarando variáveis com atribuição direta.
	name, salario := "Pit", 100

	// Chama a função setName, passando name como argumento.
	setName(name)

	// Chama a função addSalary, passando o salário e um bônus de 10.
	// Retorna dois valores: novo salário e o valor do bônus.
	newSalary, bonus := addSalary(salario, 10)

	// Exibe os valores calculados no console.
	fmt.Println("Novo salário: ", newSalary)
	fmt.Println("Bonus: ", bonus)
}

// Função para exibir o nome passado como parâmetro.
func setName(name string) {
	fmt.Println("Nome: ", name)
}

// Função para calcular o novo salário adicionando um bônus.
// Retorna o valor atualizado e o próprio bônus.
func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

// Função para retornar um nome fixo.
func getName() string {
	return "Pit"
}
