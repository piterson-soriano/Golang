package main

import "fmt"

//Channel é um canal de comunicação entre goroutines
//é um tipo seguro para concorrência
//um canal pode ser criado com a função make
//canal := make(chan tipo, buffer)
//o buffer é opcional, se não for informado o canal é bloqueante
//um canal bloqueante só permite o envio ou recebimento de um valor quando houver uma goroutine esperando
//um canal com buffer permite o envio de valores até o limite do buffer, sem bloquear a goroutine

func main() {
	ch := make(chan int, 1)
	ch <- 1 //envia o valor 1 para o canal ch
	<-ch    //recebe o valor do canal ch

	ch <- 2
	fmt.Println(<-ch)
}
