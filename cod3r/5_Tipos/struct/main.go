package main

import "fmt"

/*
Structs são tipos de dados compostos em Go que permitem agrupar diferentes tipos de dados sob um mesmo nome.
Elas são úteis para:

Organizar dados: Em vez de lidar com várias variáveis separadas, você agrupa tudo em uma única estrutura.
Representar entidades do mundo real: Como "produto", "carro", "aluno", etc.
Facilitar métodos associados: Você pode associar funções (métodos) diretamente à struct.
*/

// Define uma estrutura chamada "produto" com três campos:

type produto struct {
	nome     string
	preco    float64
	desconto float64 // desconto: valor percentual de desconto (float64, ex: 0.10 representa 10%)
}

// Define um método chamado "precoComDesconto" com receiver do tipo "produto".
// Esse método calcula o preço do produto com o desconto aplicado.
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	// Declara uma variável do tipo produto e atribui valores a seus campos
	var produto1 produto
	produto1 = produto{
		nome:     "lápis",
		preco:    1.79,
		desconto: 0.05, // 5% de desconto
	}
	// Imprime os dados do produto1 e o preço com desconto
	fmt.Println(produto1, produto1.precoComDesconto())

	// Declara e inicializa um segundo produto usando a forma curta com :=
	// e inicialização direta por posição dos campos
	produto2 := produto{"notebook", 1989.90, 0.10} // 10% de desconto
	// Imprime os dados do produto2 e o preço com desconto
	fmt.Println(produto2, produto2.precoComDesconto())
}
