package main

import "fmt"

func main() {

	// Declara variáveis com valores padrão
	var a int     // Inteiro padrão é 0
	var a1 rune   // rune é um alias para Int32
	var b float64 // Float64 padrão é 0.0
	var c bool    // Bool padrão é false
	var d string  // String padrão é ""
	var e *int    // Ponteiro para int padrão é nil

	// Exibe os valores das variáveis
	fmt.Printf("Valores iniciais:\n")
	fmt.Printf("Inteiro: %d\n", a)
	fmt.Printf("Inteiro: %d\n", a1)
	fmt.Printf("Float64: %f\n", b)
	fmt.Printf("Booleano: %t\n", c)
	fmt.Printf("String: %q\n", d)
	fmt.Printf("Ponteiro para int: %v\n\n", e)

	// Exibe os valores nulos
	fmt.Println("Valores nulos:")
	fmt.Println("Nulo para int:", a) // Inteiro nulo
	fmt.Println("Nulo para int32:", a1)
	fmt.Println("Nulo para float64:", b)      // Float64 nulo
	fmt.Println("Nulo para bool:", c)         // Bool nulo
	fmt.Println("Nulo para string:", d)       // String nula
	fmt.Println("Nulo para ponteiro int:", e) // Ponteiro nulo
}
