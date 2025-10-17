package main

import (
	"fmt"
	"time"
)

// função rotina que recebe um channel inteiro
func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Só depois 	que for lido ")
}
func main() {
	c := make(chan int) //canal sem buffer
	go rotina(c)
	fmt.Println(<-c) // operação de bloqueio
	fmt.Println("Foi lido")
	/*
		É essencial garantir que um dado seja efetivamente enviado para o canal esperado; caso contrário,
		o programa pode entrar em deadlock, ou seja, ficar bloqueado indefinidamente esperando por uma
		operação que nunca será concluída

		o que acontece na linha abaixo.
	*/
	fmt.Println(<-c)   //gera do deadlock
	fmt.Println("fim") //nunca chega aqui, pois na linha acima gera deadLock
}
