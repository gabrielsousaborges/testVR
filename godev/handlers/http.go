package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"godev/model"
	"godev/utils"
	"io"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}



func (h *Handler) CriaAluno(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Metodo incorreto")
		return
	}

	var req = &model.Aluno{}

	body, err := io.ReadAll(r.Body)
	if err != nil {

		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber o body de Transaction")
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar o unmarshal")
	}

	if err := model.InsereAluno(h.DB, req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao criar user ")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "OK")
	return

}


func (h *Handler) InsereAluno(w http.ResponseWriter, r *http.Request) {
	

	var a = &model.Aluno{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("Erro na leitura do Json de cidade %v", err)
		return
	}

	if err := json.Unmarshal(body, a); err != nil {
		fmt.Errorf("Erro no decode do Json %v", err)
		return
	}

	err = model.InsereAluno(h.DB, a)
	if err != nil {
		fmt.Errorf("Falha ao inserir aluno: %v", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "OK")
	return

}


func (h *Handler) BuscaAluno(w http.ResponseWriter, r *http.Request) {

	var (
		id   int64
		err error

	)

	if r.FormValue("ID") != "" {
		id,err = strconv.ParseInt(r.FormValue("ID"), 10, 64)
		if err != nil {
			fmt.Errorf("Erro ID")
			return
		}
	}

	aluno, err := model.BuscaAluno(h.DB, id)
	if err != nil {
		fmt.Errorf("Erro buscar Aluno")
		return
	}


	utils.RespondWithJSON(w, http.StatusOK, aluno)
	return

}