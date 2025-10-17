package main

import (
	"fmt"
	"time"
)

// funcção que envia múltiplos valores para um canal após alguns segundos
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second) //espera 1 segundo antes de enviar o valor
	c <- 2 * base           //envia 2 vezes o valor base para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int) //cria um canal de inteiros
	//chama a função em uma goroutine separada
	go doisTresQuatroVezes(2, c)
	fmt.Println("1")

	a, b := <-c, <-c  //recebendo os dois primeiros valores do canal
	fmt.Println(a, b) //imprime valores recebidos (4 e 6)

	fmt.Println(<-c) //recebe e imprime o terceiro valor do canal (8)

}
