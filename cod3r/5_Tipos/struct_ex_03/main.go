package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	p := Pessoa{
		Nome:  "Pit",
		Idade: 20,
	}
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)

	p.Idade += 1
	fmt.Println("Nova idade:", p.Idade)
}
