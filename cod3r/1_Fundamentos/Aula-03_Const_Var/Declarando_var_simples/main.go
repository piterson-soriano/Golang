package main

import "fmt"

var idade int = 35
var nome string = "Pit"
var altura float32 = 1.65

//Declaração com inferência de tipo
var cidade = "São Paulo"
var habitantes = 12000000
var telefone = 9999999999

var (
	profissao       = "Desenvolvedor"
	trabalhoHibrido = false
	anoNascimento   = 2000
)

func main() {

	//Declaração curta. Usado mais dentro de funções
	mensagem := "Olá, mundo!"
	contador := 1
	isAdmin := false

	fmt.Println("Nome: ", nome)
	fmt.Println("Idade: ", idade)
	fmt.Println("Altura: ", altura)

	fmt.Println(" ")
	fmt.Println("Cidade: ", cidade)
	fmt.Println("Habitantes: ", habitantes)
	fmt.Println("Telefone: ", telefone)

	fmt.Println(" ")
	fmt.Println("Mensagem: ", mensagem)
	fmt.Println("Contador: ", contador)
	fmt.Println("É administrador? ", isAdmin)

	fmt.Println(" ")

	fmt.Println("Profissão: ", profissao)
	fmt.Println("Trabalho presencial? ", trabalhoHibrido)
	fmt.Println("Ano de nascimento: ", anoNascimento)

}
