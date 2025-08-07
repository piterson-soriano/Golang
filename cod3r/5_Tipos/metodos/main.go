package main

import (
	"fmt"
	"strings"
)

// Define uma estrutura chamada "pessoa" com dois campos sendo, nome e sobrenome.
type pessoa struct {
	nome      string
	sobrenome string
}

// Método que retorna o nome completo da pessoa, concatenando nome e sobrenome.
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// Método que altera o nome completo da pessoa, dividindo a string em nome e sobrenome.
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	// Divide a string em partes usando espaço como separador
	partes := strings.Split(nomeCompleto, " ")
	//Atualiza os campos da estrutura com as partes obtidas
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

// Função principal que será executada ao rodar o programa
func main() {
	//Cria uma instância da estrutura pessoa com o nome "Pedro" e sobrenome "Silva"
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto()) //Imprime o nome completo

	//Altera o nome completo da pessoa para "Maria Pereira"
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto()) //Imprimi o nome completo
}
