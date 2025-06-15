package main

import "fmt"

func main() {

	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)

	fmt.Println(palavra)
	fmt.Print(invertePalavra(palavra))

}

func invertePalavra(s string) string {
	runas := []rune(s)
	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {

		runas[i], runas[j] = runas[j], runas[i]
	}

	return string(runas)

}
