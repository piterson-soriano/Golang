package main

import "fmt"

func main() {

	i := 1

	// GO não permite operações matemáticas diretamente em ponteiros
	//nil é nulo
	var p *int = nil

	p = &i //pegando o endereço da variável "i"

	// apenas "*p" para ter o valor associado ao ponteiro
	// apenas "p" para ter o endereço que está armazenado dentro do ponteiro

	*p++ //desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
