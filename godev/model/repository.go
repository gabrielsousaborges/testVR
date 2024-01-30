package model

import (
	"database/sql"
	"fmt"
	"strings"
)


func InsereAluno(conn *sql.DB, aluno *Aluno) error {
	
	defer conn.Close()
	
	err := conn.QueryRow("INSERT INTO aluno (codigo, nome) VALUES (%v, %v);", aluno.ID, aluno.Nome)
	if err != nil{
		return fmt.Errorf("falha ao criar aluno :%v", err)
	}
	return nil
}

func BuscaAluno(conn *sql.DB, nomes []string) ([]Aluno, error) {

	defer conn.Close()

	var (
		alunos []Aluno
		condicoes []string
	)
	for _, nome := range nomes {
		condicoes = append(condicoes, fmt.Sprintf("nome LIKE '%s'", nome))
	}

	resultado, err := conn.Query("SELECT * FROM aluno WHERE %s", strings.Join(condicoes, " OR "))
	if err != nil {
		return alunos, fmt.Errorf("falha na execução da busca de aluno no postgres: %v", err)
	}

	for resultado.Next() {
		var aluno Aluno
		err = resultado.Scan(&aluno.ID, &aluno.Nome)
		if err != nil {
			return alunos, err
		}
		alunos = append(alunos, aluno)
	}
	return alunos, nil
}

func AtualizaAluno(conn *sql.DB, codigo int, nome string) error {

	defer conn.Close()

	err := conn.QueryRow("UPDATE aluno SET nome=%v WHERE codigo=%v", nome, codigo)

	if err != nil {
		return fmt.Errorf("falha ao atulizar aluno:%v", err)
	}

	return nil
}

func DeletaAluno(conn *sql.DB, nome string) error {
	
	defer conn.Close()

	err := conn.QueryRow("DELETE FROM aluno WHERE nome=%v", nome)
	if err != nil {
		return fmt.Errorf("falha ao deletar aluno: %v", err)
	}

	return nil
}

func InsereCurso(conn *sql.DB, curso *Curso) error {
	
	defer conn.Close()
	
	err := conn.QueryRow("INSERT INTO curso (codigo, descricao, ementa) VALUES (%v, %v, %v);", curso.Codigo, curso.Descricao, curso.Ementa)
	if err != nil{
		return fmt.Errorf("falha ao criar curso :%v", err)
	}
	return nil
}

func BuscaCurso(conn *sql.DB, descricao []string) ([]Curso, error) {

	defer conn.Close()

	var (
		cursos []Curso
		condicoes []string
	)
	for _, descricao := range descricao {
		condicoes = append(condicoes, fmt.Sprintf("nome LIKE '%s'", descricao))
	}

	resultado, err := conn.Query("SELECT * FROM curso WHERE %s", strings.Join(condicoes, " OR "))
	if err != nil {
		return cursos, fmt.Errorf("falha na execução da busca de curso no postgres: %v", err)
	}

	for resultado.Next() {
		var curso Curso
		err = resultado.Scan(&curso.Codigo, &curso.Descricao, &curso.Ementa)
		if err != nil {
			return cursos, err
		}
		cursos = append(cursos, curso)
	}
	return cursos, nil
}

func AtualizaACurso(conn *sql.DB, codigo int, descricao string, ementa string) error {

	defer conn.Close()

	err := conn.QueryRow("UPDATE curso SET descricao=%v, ementa=%v WHERE codigo=%v", descricao, ementa, codigo)

	if err != nil {
		return fmt.Errorf("falha ao atulizar curso:%v", err)
	}

	return nil
}