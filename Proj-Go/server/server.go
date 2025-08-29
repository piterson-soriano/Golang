package main // Declara o pacote principal

import (
	"encoding/json" // Importa o pacote json para codificação e decodificação JSON
	"fmt"           // Importa o pacote fmt para formatação de saída
	"net/http"      // Importa o pacote http para criar um servidor e lidar com requisições HTTP
)

// Estrutura para a resposta JSON
type Response struct {
	Message string `json:"message"` // Campo Message, mapeado para JSON como "message"
}

// Handler para o endpoint
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Define o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")
	// Cria uma nova instância da estrutura Response com a mensagem de sucesso
	response := Response{Message: "Dados recebidos com sucesso!"}
	// Codifica a resposta como JSON e escreve no ResponseWriter
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Registra o handler para o endpoint "/api/dados"
	http.HandleFunc("/api/dados", handleRequest)
	// Exibe uma mensagem no console indicando que o servidor está rodando
	fmt.Println("Servidor rodando na porta 8080...")
	// Inicia o servidor na porta 8080 e aguarda requisições
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Se ocorrer um erro ao iniciar o servidor, exibe a mensagem de erro
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
