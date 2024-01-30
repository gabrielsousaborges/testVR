package main

import (
	"godev/configs"
	"godev/db"
	"godev/handlers"

	"github.com/gorilla/mux"
)


func main() {
	confi, err := configs.Carrega()
	if err != nil {
		panic(err)
	}


	db, err := db.AbreConexao(confi.DB)
    if err != nil {
      panic(err)
    }

	router := mux.NewRouter()

 	handlers.RegistrarRotas(router, db)
	
}
