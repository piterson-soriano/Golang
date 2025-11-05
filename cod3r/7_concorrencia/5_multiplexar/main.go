package main

import (
	"fmt"
	"time"
)

// Função que envia mensagens para um canal com um intervalo
func source(name string, interval time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s: mensagem %d", name, i)
			time.Sleep(interval)
		}
	}()
	return ch
}

// Função que multiplexa dois canais em um só
func multiplex(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case msg := <-ch1:
				out <- msg
			case msg := <-ch2:
				out <- msg
			}
		}
	}()
	return out
}

func main() {
	// Criando duas fontes de dados
	a := source("Fonte A", 1*time.Second)
	b := source("Fonte B", 2*time.Second)

	// Multiplexando as fontes
	mux := multiplex(a, b)

	// Lendo do canal multiplexado
	for i := 0; i < 10; i++ {
		fmt.Println(<-mux)
	}
}
