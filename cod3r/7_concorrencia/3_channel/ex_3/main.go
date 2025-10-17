package main

import "fmt"

//João envia o carrinho para Maria através de um túnel (canal)
func joao(tunel chan string) {
	tunel <- "Carrinho"
}

func main() {
	tunel := make(chan string) //Cria um túnel (canal) de comunicação entre goroutines
	go joao(tunel)             //João começa a brincar e envia o brinquedo pelo túnel
	brinquedo := <-tunel       //Recebe o brinquedo enviado pelo túnel
	fmt.Println("Maria recebeu o brinquedo: ", brinquedo)
}
