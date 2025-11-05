package main

import (
	"fmt"
	"time"

	"github.com/cod3rcursos/html"
)

// oMaisRapido retorna o título da primeira URL que responder
func oMaisRapido(urls ...string) string {
	// Cria um slice de canais
	canais := make([]<-chan string, len(urls))
	for i, url := range urls {
		canais[i] = html.Titulo(url)
	}

	// Canal para timeout
	timeout := time.After(1 * time.Second)

	// Multiplexa os canais usando select
	select {
	case t := <-canais[0]:
		return fmt.Sprintf("Resposta de %s: %s", urls[0], t)
	case t := <-canais[1]:
		return fmt.Sprintf("Resposta de %s: %s", urls[1], t)
	case t := <-canais[2]:
		return fmt.Sprintf("Resposta de %s: %s", urls[2], t)
	case <-timeout:
		return "Todos perderam. Time out!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://youtube.com.br",
		"https://google.com.br",
		"https://amazon.com.br",
	)
	fmt.Println("O mais rápido foi:", campeao)
}
