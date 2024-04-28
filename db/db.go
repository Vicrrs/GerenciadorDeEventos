package db

import (
  "database/sql"
  "fmt"
  "os"

  _ "github.com/lib/pq"
)

// Criando e retornando uma conexao com o banco de dados
func ConectaComBancoDeDados() *sql.DB {
  // Obtém as configuraçoes do banco de dados a partir de variáveis de ambiente
  user := os.Getenv("DB_USER")
  dbname := os.Getenv("DB_NAME")
  password := os.Getenv("DB_PASSWORD")
  host := os.Getenv("DB_HOST")
  sslmode := os.Getenv("DB_SSLMODE")

  fmt.Printf("Conectando com: user=%s dbname=%s password=%s host=%s sslmode=%s\n", user, dbname, password, host, sslmode)
  connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    panic(err.Error())
  }
  return db
}

