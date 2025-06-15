/*
Requisitos Funcionais:

1 - O programa deve iniciar perguntando ao usuário se deseja realizar um encode ou um decode.
2 - Caso o usuário opte por encode:
- Solicitar dois valores de entrada: PV e Token.
- Concatenar os valores no formato "PV:Token".
- Aplicar a codificação Base64 na string resultante.
- Exibir o valor codificado ao usuário.
3 - Caso o usuário opte por decode:
- Solicitar ao usuário um código Base64 previamente gerado.
- Realizar a decodificação do dado fornecido.
- Exibir a informação original recuperada.
4 - O programa deve operar dentro de um loop, permitindo que o usuário realize múltiplos processos de codificação e decodificação.
5 - Ao final de cada operação, perguntar ao usuário se deseja realizar um novo encode/decode.
Se a resposta for negativa, o programa deve encerrar a execução com uma mensagem final.
*/

package main

import (
	"encoding/base64" //pacote do basic
	"fmt"
)

// Função criada para codificar PV + token
func encodePVToken(pv, token string) string {
	data := pv + ":" + token                               //Realiza a concatenação
	return base64.StdEncoding.EncodeToString([]byte(data)) //Codifica a string concatenada para base64
}

// Função para decodificar
func decodeBase64(encoded string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err // Retorna erro caso a decodificação falhe
	}
	return string(decodedBytes), nil // Converte os bytes de volta para string
}

func main() {
	for {
		var escolha string
		fmt.Println("Escolha uma opção:")
		fmt.Println("1 - Fazer o encode do PV+Token")
		fmt.Println("2 - Fazer decode Base64")
		fmt.Println("3 - Sair")
		fmt.Println()
		fmt.Print("Opção: ")
		fmt.Scanln(&escolha) // Faz a leitura com a escolha feita

		switch escolha {
		case "1":
			// processo para codificar para base64
			var pv, token string
			fmt.Print("Digite o PV: ")
			fmt.Scanln(&pv) //leitura do PV
			fmt.Print("Digite o Token: ")
			fmt.Scanln(&token) //leitura do token

			//Chama a função encondePVToken para gerar um código base64
			encoded := encodePVToken(pv, token)
			fmt.Println("\nResultado do Encode:", encoded)
		case "2":
			//Codificação base64
			var encoded string
			fmt.Print("Digite o código Base64: ")
			fmt.Scanln(&encoded)

			decoded, err := decodeBase64(encoded)
			if err != nil {
				fmt.Println("Erro ao decodificar. Tente novamente!", err) //exibe erro caso falhe
			} else {
				fmt.Println("\nResultado do Decode:", decoded)
			}

		case "3":
			//Finaliza o programa.
			fmt.Println()
			fmt.Println("Programa encerrado!")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}

		fmt.Println()
	}
}
