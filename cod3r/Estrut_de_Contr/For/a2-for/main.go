package main

import "fmt"

func main() {
	soma := 0

	for i := 1; 1 <= 100; i++ {
		soma += i
	}
	fmt.Println("A soma dos números de 1 a 100 é:", soma)
}
