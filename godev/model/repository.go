package model

import (
	"database/sql"
	"fmt"
)


func InsereAluno(db *sql.DB, aluno *Aluno) error {
	

	
	defer db.Close()
	
	err := db.QueryRow("INSERT INTO alunos (ID, Nome) VALUES (%v, %v);", aluno.ID, aluno.Nome)
	if err != nil{
		return fmt.Errorf("Falha ao criar aluno")
	}
	return nil
}

func BuscaAluno(db *sql.DB, id int64) (aluno Aluno, err error) {

	defer db.Close()


	var usuario Aluno

	alunodb, err := db.Query("SELECT * FROM alunos WHERE id = %v", id)
	if err != nil {
		return Aluno{}, fmt.Errorf("falha na execução da busca de aluno no postgres: %v", err)
	}

	alunodb.Next()

	err = alunodb.Scan(&aluno.ID, &aluno.Nome)
	if err != nil {
		return Aluno{}, nil
	}
	return usuario, nil
}

func AtualizaAluno(db *sql.DB, id int64, aluno Aluno) error {

	defer db.Close()

	err := db.QueryRow("UPDATE alunos SET nome=%v", aluno.Nome)

	if err != nil {
		return fmt.Errorf("Falha ao atulizar aluno")
	}

	return nil
}

func DeletaAluno(db *sql.DB, id int64) error {
	
	defer db.Close()

	err := db.QueryRow("DELETE FROM alunos WHERE id=%v", id)
	if err != nil {
		return fmt.Errorf("Falha ao deletar aluno")
	}

	return nil
}