package main

import "fmt"

func main() {
	var idade int = 25
	const nome = "Pit"

	fmt.Println("Nome: ", nome)
	fmt.Println("Idade: ", idade)

	idade = 35
	//nome = "Maria" // erro! constante não pode mudar

	fmt.Println("Nova idade: ", idade)
}
