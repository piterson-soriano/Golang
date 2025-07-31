package main

import "fmt"

//recebe dois parâmetros trab1 e trab2 do tipo bool
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2 //&& operador "E"
	comprarTv32 := trab1 != trab2 //ou operador exclusivo/XOR. Pode ser comprada, se um dos parâmetros for true
	//operador OR (||), que retorna true se pelo menos um dos parâmetros for true. O sorvete pode ser comprado se pelo menos um dos trabalhos for realizado.
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	//Chamando a função compras passando true para os dois parâmetros trab1 e trab2.
	//Na função compras retorna 3 valores booleanos tv50, tv32, sorvete
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf(" Tv50: %t\n Tv32: %t\n Sorvete: %t\n Saudável: %t\n",
		tv50, tv32, sorvete, !sorvete)

}
