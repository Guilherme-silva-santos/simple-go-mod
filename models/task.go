package models

// criando a tipagem para as models de task
// Utilizamos essas tags json:"..." para informar ao Go que,
// quando essa struct for convertida para JSON — por exemplo,
// em uma requisição GET ou POST — os campos do JSON devem ter os
// nomes definidos após os dois-pontos (:).
// E como no go os campos das struct precisam ser com letra maiúscula,
// podemos fazer desse jeito para que os campos do json sigam um outro padrão.
// basicamente quando for feito o encode para jason os campos ficaram assim
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// criando tabela para o banco de dados
// migration
const (
	TableName   = "tasks"
	CreateTable = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT NOT NULL,
		status BOOLEAN NOT NULL DEFAULT FALSE
)
`
)
