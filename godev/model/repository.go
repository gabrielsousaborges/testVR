package model

import (
	"fmt"
	"godev/db"
)


func InsereAluno(aluno *Aluno) error {
	
	conn := db.AbreConexao()
	
	defer conn.Close()
	
	err := conn.QueryRow("INSERT INTO alunos (ID, Nome) VALUES (%v, %v);", aluno.ID, aluno.Nome)
	if err != nil{
		return fmt.Errorf("Falha ao criar aluno")
	}
	return nil
}

func BuscaAluno(id int64) (aluno Aluno, err error) {
	conn:= db.AbreConexao()

	var usuario Aluno

	alunodb, err := conn.Query("SELECT * FROM alunos WHERE id = %v", id)
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

func AtualizaAluno(id int64, aluno Aluno) error {

	conn := db.AbreConexao()

	defer conn.Close()

	err := conn.QueryRow("UPDATE alunos SET nome=%v", aluno.Nome)

	if err != nil {
		return fmt.Errorf("Falha ao atulizar aluno")
	}

	return nil
}

func DeletaAluno(id int64) error {
	conn := db.AbreConexao()

	
	defer conn.Close()
	
	err := conn.QueryRow("DELETE FROM alunos WHERE id=%v", id)
	if err != nil {
		return fmt.Errorf("Falha ao deletar aluno")
	}

	return nil
}