package main

import "fmt"

func rotina(ch chan int) { // rotina envia 6 valores para o canal
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou!") // só executa quando há espaço no buffer
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // canal com buffer de 3
	go rotina(ch)           // rotina envia 6 valores para o canal

	fmt.Println(<-ch) // recebe 3 valores do canal
}
