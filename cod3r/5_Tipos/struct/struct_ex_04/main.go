package main

import "fmt"

type Carro struct {
	Modelo string
	Ano    int
}

func (c Carro) Descricao() string {
	return fmt.Sprintf("Carro modelo %s do ano %d ", c.Modelo, c.Ano)
}
func main() {
	c := Carro{
		Modelo: "esportivo",
		Ano:    2025,
	}
	c = Carro{"VOLKSWAGEN FUSCA 1300", 1970}
	fmt.Println("Descrição - ", c.Descricao())
	fmt.Println("O modelo do carro é ", c.Modelo+" e do ano de ", c.Ano)

}
