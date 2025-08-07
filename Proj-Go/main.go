package main // Declara o pacote principal

import (
	"bytes"         // Importa o pacote bytes para manipulação de bytes
	"encoding/csv"  // Importa o pacote csv para leitura de arquivos CSV
	"encoding/json" // Importa o pacote json para codificação e decodificação JSON
	"fmt"           // Importa o pacote fmt para formatação de saída
	"net/http"      // Importa o pacote http para fazer requisições HTTP
	"os"            // Importa o pacote os para interagir com o sistema operacional
)

// Estrutura p/ os dados a serem enviados
// Campo Nome, mapeado para JSON como "nome"
// Campo Email, mapeado para JSON como "email"
type Dados struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	// Abre o arquivo CSV
	file, err := os.Open("dados.csv") // Tenta abrir o arquivo "dados.csv"
	if err != nil {                   // Verifica se houve erro ao abrir o arquivo
		fmt.Println("Erro ao abrir o arquivo:", err) // Exibe a mensagem de erro
		return                                       // Encerra a função main se houver erro
	}
	defer file.Close() // Garante que o arquivo será fechado ao final da função

	// Lê o conteúdo do arquivo CSV
	reader := csv.NewReader(file)    // Cria um novo leitor CSV para o arquivo
	records, err := reader.ReadAll() // Lê todos os registros do arquivo
	if err != nil {                  // Verifica se houve erro ao ler o arquivo
		fmt.Println("Erro ao ler o arquivo CSV:", err) // Exibe a mensagem de erro
		return                                         // Encerra a função main se houver erro
	}

	// Itera sobre os registros
	for _, record := range records[1:] { // Ignora o cabeçalho, começando a iteração na segunda linha
		dados := Dados{ // Cria uma nova instância da estrutura Dados
			Nome:  record[0], // Atribui o primeiro campo da linha ao campo Nome
			Email: record[1], // Atribui o segundo campo da linha ao campo Email
		}

		// Converte os dados para JSON
		jsonData, err := json.Marshal(dados) // Converte a estrutura Dados para formato JSON
		if err != nil {                      // Verifica se houve erro na conversão
			fmt.Println("Erro ao converter dados para JSON:", err) // Exibe a mensagem de erro
			continue                                               // Continua para a próxima iteração do loop
		}

		// Envia a requisição POST para o endpoint
		resp, err := http.Post("http://localhost:8080/api/dados", "application/json", bytes.NewBuffer(jsonData)) // Envia a requisição POST
		if err != nil {                                                                                          // Verifica se houve erro ao enviar a requisição
			fmt.Printf("Erro ao fazer a requisição para %s: %v\n", dados.Email, err) // Exibe a mensagem de erro
			continue                                                                 // Continua para a próxima iteração do loop
		}
		defer resp.Body.Close() // Garante que o corpo da resposta será fechado após o uso

		// Exibe o resultado da requisição
		fmt.Printf("Requisição feita para %s (%s) - Status: %s\n", dados.Nome, dados.Email, resp.Status) // Mostra o nome, email e status da resposta
	}
}
