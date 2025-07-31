package main

import "fmt"

/*
Usar %s!\n permite que você crie mensagens dinâmicas, substituindo o %s pelo valor desejado
e garantindo que a saída seja formatada de forma clara e legível no console.


%s: Indica onde colocar uma string. O %s foi substituído pela string "Pit" "." "Go"
\n: A nova linha ajuda a separar as saídas, se necessário.
*/
func main() {
	fmt.Printf("Meu nome é %s\n", "Pit")
	fmt.Printf("Outro programa em %s %s %s %s", "Go", ".", ".", ".")

	fmt.Printf(" Teste %s %s \n", "-", "com printf")

	fmt.Printf("Imprimindo com %s %s \n ", "comando", "printf")

}
