package main

import "fmt"

//interface em Go, é como um contrato.

//Qualquer coisa que tenha um método chamado toString que retorna uma string, pode ser tratada como um imprimivel.
type pit interface { // Declara uma interface chamada "imprimivel" com um método "toString" que retorna uma string.
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}
type produto struct {
	nome  string
	preco float64
}

// Implementação do método "toString" para o tipo "pessoa".
// Interfaces em Go são implementadas implicitamente: basta que o tipo tenha os métodos exigidos.
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome // Retorna o nome completo da pessoa.
}

// Implementação do método "toString" para o tipo "produto".
func (p produto) toString() string {
	// Usa fmt.Sprintf para formatar a string com nome e preço com duas casas decimais.
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// Função que recebe qualquer tipo que implemente a interface "imprimivel" e imprime seu conteúdo.
func imprimir(x pit) {
	fmt.Println(x.toString()) // Chama o método "toString" e imprime o resultado
}

func main() {
	// Cria uma variável do tipo "imprimivel" e atribui uma instância de "pessoa".
	var coisa pit = pessoa{"Roberto", "Silva"}

	// Imprime diretamente o resultado de toString().
	fmt.Println(coisa.toString())

	// Chama a função "imprimir" passando a pessoa
	imprimir(coisa)

	// Reatribui a variável "coisa" com uma instância de "produto".
	coisa = produto{"Calça jeans", 78.99}

	// Imprime diretamente o resultado de toString()
	fmt.Println(coisa.toString())

	// Chama a função "imprimir" passando o produto.
	imprimir(coisa)

}
