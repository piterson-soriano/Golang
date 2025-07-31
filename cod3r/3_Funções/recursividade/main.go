package main

import "fmt"

//função chamada fatorial que recebe um número inteiro n. Retorna dois valores: Um inteiro com o resultado. Um erro, caso algo dê errado
func fatorial(n int) (int, error) {
	switch
	//Se o número for menor que 0, retorna um erro dizendo que o número é inválido — pois não existe fatorial de número negativo.
	{
	// Caso o número seja menor que 0, retornamos um erro, pois não existe fatorial de número negativo.
	case n < 0:
		return -1,
			fmt.Errorf("Número inválido: %d.", n)
		// Caso base: o fatorial de 0 é 1
	case n == 0:
		return 1,
			nil // nil indica que não há erro
	default:
		// Chamamos a função novamente com n - 1
		// O segundo valor retornado (erro) é ignorado usando "_", pois já validamos que n é positivo
		fatorialAnterior, _ := fatorial(n - 1) // Chamada recursiva
		return n * fatorialAnterior, nil       // Retorna o fatorial do número atual
	}
}

func main() {
	// Calcula o fatorial de 5 e ignora o erro retornado
	resultado, _ := fatorial(5)
	// Exibe o resultado do fatorial de 5
	fmt.Println(resultado)
	// Tenta calcular o fatorial de um número negativo
	_, err := fatorial(-4)
	// Verifica se houve erro e exibe a mensagem de erro no terminal
	if err != nil {
		fmt.Println(err)
	}
}
