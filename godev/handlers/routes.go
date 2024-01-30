package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)


func RegistrarRotas(router *mux.Router, db *sql.DB) {
      
	h := &Handler{
	  DB: db,
	}
	router.HandleFunc("/aluno/resgistrar", h.InsereAluno)
	router.HandleFunc("/aluno/criar", h.CriaAluno)
	router.HandleFunc("/aluno/buscar", h.BuscaAluno)
  }