package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilherme-silva-santos/simple-go-mod/config"
	"github.com/guilherme-silva-santos/simple-go-mod/handlers"
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

	defer dbConnection.Close()
	if err != nil {
		log.Fatal("Error creating table")
	}

	// end-points
	router := mux.NewRouter()

	// aqui estamos criando um end-point para o método GET
	// passando para ele o path de task

	taskHandler := handlers.NewTaskHandler(dbConnection)
	// na rota de tasks ele vai a instacia de conexão com o banco
	// passando o metodo de leitura de tasks
	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTasks).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTasks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
