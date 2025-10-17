package main

import (
	"fmt"
	"time"
)

/*
As goroutines são úteis na programação porque elas ajudam o computador a fazer várias coisas ao mesmo tempo, de forma rápida,
leve e organizada. Aqui vai um resumo simples e direto:

Por que usar goroutines?
- Concorrência eficiente: Você pode executar várias tarefas ao mesmo tempo sem precisar esperar uma terminar para começar outra.
- Baixo custo: Diferente das threads tradicionais, goroutines usam pouquíssima memória e são muito rápidas para iniciar.
- Alta escalabilidade: É possível criar milhares (ou milhões!) de goroutines sem travar o sistema.
- Facilidade de uso: Basta colocar a palavra go antes da função e pronto — ela roda em paralelo.
- Comunicação segura: Com os canais (channels), você consegue fazer goroutines conversarem entre si sem bagunça ou
 erros de sincronização.

Exemplos de uso
- Servidores web que atendem muitos usuários ao mesmo tempo.
- Processamento de arquivos ou dados em paralelo.
- Aplicativos que precisam fazer várias tarefas sem travar a interface.


*/

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d) \n", pessoa, texto, i+1)
	}

}
func main() {
	//fale("Maria", "Pq vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	//go fale("Maria", "Ei...", 500) //goroutine
	//go fale("João", "opa...", 500) //goroutine

	go fale("Maria", "ei...", 10) //goroutine
	fale("João", "opa...", 5)

}
