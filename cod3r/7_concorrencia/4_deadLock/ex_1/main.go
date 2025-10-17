package main

func main() {

	//ex de canal com buffer
	//ch := make(chan int, 1) //criando um canal com buffer de 1 posição
	//ch <- 2                 //não entra em deadlock
	ch := make(chan int) //criando um canal sem buffer
	ch <- 2              // Deadlock aqui!
}
