package handlers

import (
	"encoding/json"
	"fmt"
	"godev/configs"
	"godev/model"
	"io"
	"net/http"
	"strconv"
)

func CriaAluno(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		configs.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Metodo incorreto")
		return
	}

	var req = &model.Aluno{}

	body, err := io.ReadAll(r.Body)
	if err != nil {

		configs.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber o body de Transaction")
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		configs.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao realizar o unmarshal")
	}

	if err := model.InsereAluno(req); err != nil {
		configs.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao criar user ")
		return
	}

	configs.RespondWithJSON(w, http.StatusOK, "OK")
	return

}

func AlunoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		BuscaAluno(w, r)
		return
	}

	if r.Method == http.MethodPost {
		InsereAluno(w, r)
		return
	}

	fmt.Errorf("Erro metodo")
	return
}

func InsereAluno(w http.ResponseWriter, r *http.Request) {
	

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

	err = model.InsereAluno(a)
	if err != nil {
		fmt.Errorf("Falha ao inserir aluno: %v", err)
		return
	}

	configs.RespondWithJSON(w, http.StatusOK, "OK")
	return

}


func BuscaAluno(w http.ResponseWriter, r *http.Request) {

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

	aluno, err := model.BuscaAluno(id)
	if err != nil {
		fmt.Errorf("Erro buscar Aluno")
		return
	}


	configs.RespondWithJSON(w, http.StatusOK, aluno)
	return

}