package main

import (
	"math/rand"
	"time"
)

func numroAleatorio() int {
	// rand - pacote na biblioteca padrão do Go que fornece funções para geração de números aleatórios
	//NewSource é uma função desse pacote que cria uma nova fonte de números aleatórios.
	s := rand.NewSource(time.Now().UnixNano()) //declara e inicializa uma nova chamada "s"
	r := rand.New(s)
}
