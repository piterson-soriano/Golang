package main

import "fmt"

//Criei uma função chamada novaTransacao que recebe um identificador de transação (ex: "TX123").
// Essa função devolve outra função como resposta.
func novaTransacao(id string) func(string) {
	status := "pendente" //status inicial da transação

	return func(novoStatus string) { //	novoStatus é o status que será atualizado
		fmt.Printf("Transação %s: %s ➡ %s\n", id, status, novoStatus) // imprime o status atual e o novo status
		status = novoStatus                                           //atualiza o status da transação

	}
}

func main() {
	atualiza := novaTransacao("TX123") // Cria uma nova transação com o ID "TX123" e obtém a função de atualização

	atualiza(" pendente")  // Atualiza o status da transação para "pendente"
	atualiza(" capturada") // Atualiza o status da transação para "capturada"
	atualiza(" cancelada") // Atualiza o status da transação para "cancelada"
}
