package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// como o pq é um driver para o banco de dados postgres
// é utilizado o _ pq não vamos usar ele diretamente

// sql.db pois essa função vai retornar uma instancia do meu db
// foi usando o ponteiro pois queremos que a função
// retorne o endereço de momeoria
// da variavel que guardou a minha conexão com o banco de dados
// e com isso será possivel usar essa conexão com todos
// os endpoins, ao invez de abrir uma nova conexão para cada end-point
func SetupDb() *sql.DB {
	// função setupDb inicializa a conexão com o banco de dados
	// inicia as variaveis de ambiente
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// sprintf formata uma string e retorna ela
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(connectionString)
	// abre a conexão com o banco de dados
	dbConnection, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Error opening database connection")
	}

	// verifica se a conexão com o banco de dados está ok
	// e como ja haviamos declarado o erro antes
	// agora apenas o err recebe o oq o ping retorna
	// não precisamos := atribuir o oq retornar a err
	err = dbConnection.Ping()

	if err != nil {
		log.Fatal("Error pinging database")
	}
	fmt.Println("Connected to database")

	return dbConnection
}
