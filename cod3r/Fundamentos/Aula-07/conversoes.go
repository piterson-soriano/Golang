package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinalInt := int(nota)
	notaFinalFloat := float32(nota)

	fmt.Println("Nota Int:", notaFinalInt)
	fmt.Println("Nota Float:", notaFinalFloat)

	//concatenar string com valor int
	fmt.Println("Teste " + strconv.Itoa(123))

	/*
		Conversão de String para Inteiro

		Esta função do pacote strconv converte uma string para um inteiro. O nome "Atoi" significa "ASCII to integer".
		O parâmetro "123" é a string que estamos tentando converter.
		A função atoi retorna dois valores (num e _)
		O uso de _ ignora o segundo valor, que é o erro. Isso é comum quando você tem certeza de que a conversão não falhará.
	*/
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	fmt.Println("-----------------Condição-----------------")

	//b - Esta função do pacote strconv tenta converter a string "1" em um valor booleano. No contexto do Go, "1" é considerado equivalente a true.
	//err - um valor de erro que indica se a conversão foi bem-sucedida. Se a conversão falhar, err conterá informações sobre o erro.
	b, err := strconv.ParseBool("")

	//Este bloco verifica se ocorreu um erro durante a conversão.
	//Se err não for nil, significa que a conversão falhou
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		//encerra a execução da função atual, evitando que a próxima instrução seja executada, já que o erro foi encontrado.
		return
	}

	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("falso")
	}
}
