package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite.")
	}

	// Formata a data e hora no formato "DD/MM/YYYY HH:mm:ss"
	// Utiliza o template de data "02/01/2006" e hora "15:04:05"
	fmt.Println("Hora atual: ", t.Format("02/01/2006 15:04:05"))
}
