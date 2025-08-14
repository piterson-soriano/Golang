package main

import "fmt"

type bike struct {
	marca         string
	preco         float64
	tipo          string
	velocidadeMax int
}

type motorizada struct {
	bike
	bikeEletrica bool
}

func main() {
	b := motorizada{}
	b.marca = "caloi"
	b.velocidadeMax = 80
	b.preco = 10000.00

	fmt.Printf("A bike da marca %s consegue atingir a velocidade de %d km/h\n", b.marca, b.velocidadeMax)

}
