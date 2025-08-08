package main

import "fmt"

//Struct chamada carro
type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //Na struct Ferrari atribui os campos/atributos do carro
	turboLigado bool
}

func main() {
	f := ferrari{} //Variavel do tipo ferrari
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true
	fmt.Printf("A ferrari %s está com turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
