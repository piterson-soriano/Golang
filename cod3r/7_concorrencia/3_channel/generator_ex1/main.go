package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			defer func() {
				if r := recover(); r != nil {
					c <- fmt.Sprintf("Erro ao processar %s", url)
				}
			}()

			resp, err := http.Get(url)
			if err != nil {
				c <- fmt.Sprintf("Erro ao acessar %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("Erro ao ler %s: %v", url, err)
				return
			}

			r := regexp.MustCompile("(?i)<title>(.*?)</title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) >= 2 {
				c <- matches[1]
			} else {
				c <- fmt.Sprintf("Título não encontrado em %s", url)
			}
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.youtube.com", "https://www.amazon.com")

	fmt.Println("Primeiros: ", <-t1, " | ", <-t2)
	fmt.Println("Segundos: ", <-t1, " | ", <-t2)
}
