package main

import "fmt"

//Variaveis da estrutura item
type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

//nome da função: valoTotal
//responsavel por calcular o valor total do pedido
func (p pedido) valorTotal() float64 {
	//variavel
	total := 0.0
	//slice
	// _, é usado para ignorar o valor do item
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.19},
		},
	}
	fmt.Printf("Valor total do pedido: %.2f ", pedido.valorTotal())
}
