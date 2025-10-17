package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		fmt.Println(<-c) //Recebe o valor do canal
	}()

	c <- 1 //sem essa linha o programa entra em deadlock
}
