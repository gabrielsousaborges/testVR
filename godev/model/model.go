package model

type Aluno struct {

	ID int64 `json: id`
	Nome string `json: nome`

}

type Curso struct {
	Codigo int64
	Descricao string
	emenda string
	Alunos []Aluno
}

type CursoAluno struct {
	Codigo int64
	CodigoAluno int64
	CodigoCurso int64
}