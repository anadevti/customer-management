package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é uma variável global para a conexão com o banco de dados
var DB *gorm.DB

// ConnectDatabase é uma função exportada que inicializa a conexão com o banco de dados, neste caso o PostgreSQL
func ConnectDatabase() {
	// Carregando o .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Para Obter variáveis de ambiente
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Montar o DSN
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=America/Sao_Paulo"

	// Conectar ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	DB = db // Atribuir a conexão à variável global DB
	log.Println("Conectado ao banco de dados com sucesso!")
}
