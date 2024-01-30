package handlers

import (
	"database/sql"
	"encoding/json"
	"godev/domain/aluno"
	"godev/model"
	"godev/utils"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}


func (h *Handler) InsereAluno(w http.ResponseWriter, r *http.Request) {
	

	var a = &model.Aluno{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro na leitura do Json de aluno %v")
		return
	}

	if err := json.Unmarshal(body, a); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro no decode do Json")
		return
	}

	err = model.InsereAluno(h.DB, a)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Falha ao inserir aluno")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "OK")
	

}


func (h *Handler) BuscaAluno(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.FormValue("nome") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Necessario informar um nome")
		return
	}

	nomes := aluno.NormalizaNome(r.FormValue("nome"))

	alunos, err := model.BuscaAluno(h.DB, nomes)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro buscar Aluno")
		return
	}


	utils.RespondWithJSON(w, http.StatusOK, alunos)

}

func (h *Handler) AtualizaAluno(w http.ResponseWriter, r *http.Request) {

	var (
		codigo int
		err error
	)

	if r.FormValue("nome") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Necessario informar um nome")
		return
	}

	nome := r.FormValue("nome")

	if r.FormValue("codigo") != "" {
		codigo, err = strconv.Atoi(r.FormValue("codigo"))
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro ao receber codigo")
			return
		}
	}


	err = model.AtualizaAluno(h.DB, codigo, nome)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro Atualizar Aluno")
		return
	}


	utils.RespondWithJSON(w, http.StatusOK, "Atualizado")

}


func (h *Handler) InsereCurso(w http.ResponseWriter, r *http.Request) {
	

	var c = &model.Curso{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro na leitura do Json de curso %v")
		return
	}

	if err := json.Unmarshal(body, c); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro no decode do Json")
		return
	}

	err = model.InsereCurso(h.DB, c)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Falha ao inserir curso")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "OK")
	

}


func (h *Handler) BuscaCurso(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.FormValue("descricao") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Necessario informar um nome")
		return
	}

	desc := aluno.NormalizaNome(r.FormValue("descricao"))
	curso, err := model.BuscaCurso(h.DB, desc)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro buscar Aluno")
		return
	}


	utils.RespondWithJSON(w, http.StatusOK, curso)

}

func (h *Handler) AtualizaCurso(w http.ResponseWriter, r *http.Request) {

	var (
		codigo int
		err error
	)

	if r.FormValue("descricao") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Necessario informar uma descricao")
		return
	}

	descricao := r.FormValue("descricao")

	if r.FormValue("ementa") == "" {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Necessario informar uma ementa")
		return
	}

	ementa := r.FormValue("ementa")

	if r.FormValue("codigo") != "" {
		codigo, err = strconv.Atoi(r.FormValue("codigo"))
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro ao receber codigo")
			return
		}
	}


	err = model.AtualizaACurso(h.DB, codigo, descricao, ementa)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0 , "Erro Atualizar Aluno")
		return
	}


	utils.RespondWithJSON(w, http.StatusOK, "Atualizado")

}