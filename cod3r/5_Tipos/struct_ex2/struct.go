package main

import "fmt"

// Uma struct em Go serve para agrupar dados relacionados em um único tipo composto.
// É como um modelo ou “molde” para criar objetos com vários campos, cada um contendo um valor que pode ser de tipos diferentes.
type carro struct {
	marca      string
	modelo     string
	preco      float64
	ano        int
	cor        string
	automatico bool
	portas     int
	placa      string
}

func (c carro) precoComDesconto() float64 {
	return c.preco * (1 - 0.18)

}

func main() {
	var carro1 carro
	carro1 = carro{
		//Atributos do carro
		marca:      "Toyota",
		modelo:     "Corolla Cross 2025",
		preco:      164000.00,
		ano:        2025,      // exemplo de valor
		cor:        "Grafite", // exemplo de valor
		automatico: true,      // exemplo de valor
		portas:     4,         // exemplo de valor
		placa:      "ABC1D23", // exemplo de valor
	}
	fmt.Printf("Marca: %s\n", carro1.marca)
	fmt.Printf("Modelo: %s\n", carro1.modelo)
	fmt.Printf("Ano: %d\n", carro1.ano)
	fmt.Printf("Cor: %s\n", carro1.cor)
	fmt.Printf("Automático: %t\n", carro1.automatico)
	fmt.Printf("Portas: %d\n", carro1.portas)
	fmt.Printf("placa: %s\n", carro1.placa)
	fmt.Printf("Preço: R$ %.2f\n", carro1.preco)
	fmt.Printf("Preço com desconto: R$ %.2f\n", carro1.precoComDesconto())

	//fmt.Println(carro1, carro1.precoComDesconto())
}
