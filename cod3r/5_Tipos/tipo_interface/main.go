package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println("Impressão interface: ", coisa)

	coisa = 3
	fmt.Println("Impressão valor: ", coisa)

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println("Impressão mensagem: ", coisa2)

	coisa2 = true
	fmt.Println("Impressão bool: ", coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}
