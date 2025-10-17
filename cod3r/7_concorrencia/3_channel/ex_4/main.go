package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "olá"
	ch <- "mundo"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
