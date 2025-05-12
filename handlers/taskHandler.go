package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guilherme-silva-santos/simple-go-mod/models"
)

type TaskHandler struct {
	// o db precisa ser um ponteiro para o tipo
	// da instancia da conexão do banco de dados
	DB *sql.DB
}

// função recebe um ponteiro para conxão
// e salva isso salva isso numa instancia de taskHandler
// e retorna retorna o endereço dessa instancia
func NewTaskHandler(db *sql.DB) *TaskHandler {
	// meu taskhandler vai acessar o endreço de momoria do db

	// salvando o db em DB e retorna o endereço de momoria dessa instancia
	// e com isso podemos usar o db em todos os end-points
	return &TaskHandler{DB: db}
}

// metodos do taskHandler

// a func recebe um task handler que vai ser um ponteiro para o
// meu tipo e nesse tipo tem a minha instancia d econexão com com o banco
func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	// aqui vamos fazer a chamada para o banco de dados

	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks []models.Task

	// serve para iterar sobre as linhas que foram retornadas
	for rows.Next() {
		var task models.Task
		// scan le todas as informações daquela linha do banco de dados
		// e salva essa informações em um ponteiro tambem.
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// o append adiciona o task que foi lido do banco de dados
		// oq ele az como usamos o next para pegar linha por linha
		// ele tem um array das task iniciando como vazio e a cada
		// linha retornada ele adiciona uma task nova nesse array
		tasks = append(tasks, task)
	}
	// usamos o writter para retornar a resposta para o cliente
	// que sera do tipo json
	w.Header().Set("Content-Type", "application/json")
	// usa o json fala para ele escrever direto na resposta
	// e encode converte o array de task para o formato json
	json.NewEncoder(w).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	// aqui ele vai pegar o body da requisição e vai fazer a leitura
	// e vai fazer o decode para a struct de task
	// e se não der erro ele vai salvar a task no endereço de momoria
	// de task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// aqui ele faz a inserção no banco de dados
	err = taskHandler.DB.QueryRow("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id", task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (taskHandler *TaskHandler) UpdateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var task []models.Task
	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := taskHandler.DB.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", task[0].Title, task[0].Description, task[0].Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsRowsAffected, err := results.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsRowsAffected == 0 {
		http.Error(w, "No rows affected", http.StatusNotFound)
		return

	}

	task[0].ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task[0])
}

func (taskHandler *TaskHandler) DeleteTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No task found with this id", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
