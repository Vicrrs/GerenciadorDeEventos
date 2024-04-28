package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Vicrrs/GerenciadorDeEventos/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Definindo a porta do servidor como variável de ambiente ou usa um padrao
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão se não estiver definida em .env
	}

	// Carrega as rotas definidas no pacote routes
	routes.CarregaRotas()

	// Inicia o servidor na porta especificada, ouve o server e serve as requisições HTTP
	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
