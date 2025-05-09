package main

import (
	"log"
	"net/http"

	"github.com/guilherme-silva-santos/simple-go-mod/config"
	"github.com/guilherme-silva-santos/simple-go-mod/models"
)

func main() {
	dbConnection := config.SetupDb()
	// defer serve para quando a função main para de executar(quando
	// o programa terminar ou dar erro), ele faz a chamada para
	// fechar a conexão com o banco de dados.
	// basicamente ele encerra algo quando a função termina
	// nesse caso, ele vai fechar a conexão com o banco de dados

	// ignora a resposta da crição da tabela com o blanck
	// e pega apenas o erro
	// e como esta criando a tabela ?
	// pq a nossa string de criação da tabela e usa o metodo exec, para
	// para que execute no banco que está conectado
	// ele fala pegue a string SQL armazenada em models
	// e execute ela no banco de dados usando a conexão.
	_, err := dbConnection.Exec(models.CreateTable)

	if err != nil {
		log.Fatal("Error creating table")
	}

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
