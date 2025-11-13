package matematica

import (
	"fmt"
	"strconv"
)

// media é função que calcula a média aritmética dos números fornecidos.
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	/*
		1. fmt.Sprintf("%.2f", media)
		- Formata o valor de media como uma string com duas casas decimais.

		2. strconv.ParseFloat(..., 64)
		- Converte a string formatada de volta para o tipo float64.

		- O segundo valor retornado (um possível erro) é ignorado usando _.

	*/

	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada

	/*
				linha de comando:
				go test -coverprofile=resultado.out
		- Gera um relatório de cobertura de testes (quais partes do código foram testadas).


				go tool cover -html=resultado.out
		- Exibe um relatório visual em HTML mostrando a cobertura de código.

	*/
}
