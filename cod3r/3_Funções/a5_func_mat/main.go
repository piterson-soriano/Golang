package main

import "fmt"

func somar(a, b float64) float64 {
	return a + b
}

func subtrair(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Atenção! Não é possível dividir por zero!")
		return 0
	}
	return a / b
}

func main() {
	num1 := 10.0
	num2 := 2.0

	fmt.Println("Soma: ", somar(num1, num2))
	fmt.Println("Subtrair: ", subtrair(num1, num2))
	fmt.Println("Multiplicar: ", multiplicar(num1, num2))
	fmt.Println("Dividir: ", dividir(num1, num2))
}
