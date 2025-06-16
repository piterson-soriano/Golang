package main

import (
	"encoding/base64"
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
