package main

import "fmt"

func main() {
	var b byte = 3 // Declara uma variável 'b' do tipo byte e atribui o valor 3
	fmt.Println(b) // Imprime o valor de 'b'

	i := 3 // Declara a variável 'i' e inicializa com 3 (inferência de tipo)
	i += 3 // Adiciona 3 a 'i' (equivalente a i = i + 3)
	i -= 3 // Subtrai 3 de 'i' (equivalente a i = i - 3)
	i /= 2 // Divide 'i' por 2 (equivalente a i = i / 2)
	i *= 2 // Multiplica 'i' por 2 (equivalente a i = i * 2)
	i %= 2 // Calcula o resto da divisão de 'i' por 2 (equivalente a i = i % 2)

	fmt.Println(i) // Imprime o valor final de 'i'

	x, y := 1, 2      // Declara as variáveis 'x' e 'y' e inicializa com 1 e 2, respectivamente
	fmt.Println(x, y) // Imprime os valores de 'x' e 'y'

	x, y = y, x       // Troca os valores de 'x' e 'y' (agora x é 2 e y é 1)
	fmt.Println(x, y) // Imprime os novos valores de 'x' e 'y'
}
