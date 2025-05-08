package config

import (
	"database/sql"
	"fmt"
)

// sql.db pois essa função vai retornar uma instancia do meu db
// foi usando o ponteiro pois queremos que a função
// retorne o endereço de momeoria
// da variavel que guardou a minha conexão com o banco de dados
// e com isso será possivel usar essa conexão com todos
// os endpoins, ao invez de abrir uma nova conexão para cada end-point
func SetupDb() *sql.DB {
	// sprintf formata uma string e retorna ela
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")
}
