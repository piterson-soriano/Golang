package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 //envia o valor 1 para o canal ch
	<-ch    //recebe o valor do canal ch

	ch <- 2
	fmt.Println(<-ch)
}
