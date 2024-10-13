package main

import (
	"fmt"
	f "fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	//forma reduzida de ciar uma var
	area := PI * m.Pow(raio, 2)
	f.Println("A área da circunsferência é ", area)

	//bloco de construção de constantes
	const (
		a = 1
		b = 2
	)
	//bloco de construção de variaveis
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)
	//Declarando na mesma linha várias variáveis
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
