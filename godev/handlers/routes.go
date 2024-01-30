package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)


func RegistrarRotas(router *mux.Router, conn *sql.DB) {
      
	h := &Handler{
	  DB: conn,
	}
	router.HandleFunc("/aluno", h.InsereAluno).Methods("POST")
	router.HandleFunc("/aluno", h.AtualizaAluno).Methods("PUT")
	router.HandleFunc("/aluno", h.BuscaAluno).Methods("GET")
	router.HandleFunc("/curso", h.InsereCurso).Methods("POST")
	router.HandleFunc("/curso", h.AtualizaCurso).Methods("PUT")
	router.HandleFunc("/curso", h.BuscaCurso).Methods("GET")
  }