package main

import (
	"godev/configs"
	"godev/handlers"

	"github.com/gorilla/mux"
)


func main() {
	err := configs.Carrega()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	 
	router.HandleFunc("/aluno/resgistrar", handlers.InsereAluno)

	
}
