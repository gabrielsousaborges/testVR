package db

import "database/sql"

func CriaTabela(conn *sql.DB) error {

	queries := []string {
		`CREATE TABLE IF NOT EXISTS aluno (codigo int PRIMARY KEY, nome varchar(50))`,
		`CREATE TABLE IF NOT EXISTS curso (codigo int PRIMARY KEY, descricao varchar(50), ementa text)`,
		`CREATE TABLE IF NOT EXISTS curso_aluno (codigo int PRIMARY KEY, codigo_aluno int, codigo_curso int, FOREIGN KEY codigo_aluno REFERENCES aluno(codigo), FOREIGN KEY codigo_curso REFERENCES curso(codigo))`,
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	} ()
	for _, query := range queries {
		_, err = tx.Exec(query)
		if err != nil {
			return err
		}
	}
		return nil
}