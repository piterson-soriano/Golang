package main

import "fmt"

func main() {

	var name, secondName string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&name)
	fmt.Println("--------------------")
	fmt.Print("Sobrenome: ")
	fmt.Scan(&secondName)
	fmt.Println("--------------------")
	fmt.Println("Nome compleo: " + name + " " + secondName)

	runas := []rune(name + " " + secondName)

	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}
	nameInvertido := string(runas)
	fmt.Println("Nome invertido: ", nameInvertido)
}
