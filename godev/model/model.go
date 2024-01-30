package model

type Aluno struct {
	ID int64 `json: codigo`
	Nome string `json: nome`

}

type Curso struct {
	Codigo int64 `json: codigo`
	Descricao string `json: descricao`
	Ementa string `json: ementa`
	Alunos []Aluno `json: alunos`
}

type CursoAluno struct {
	Codigo int64 `json: codigo`
	CodigoAluno int64 `json: codigo_aluno`
	CodigoCurso int64 `json: codigo_curso`
}