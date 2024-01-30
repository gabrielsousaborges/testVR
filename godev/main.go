package main

import (
	"godev/configs"
	"godev/db"
	"godev/handlers"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	confi, err := configs.Carrega()
	if err != nil {
		panic(err)
	}


	conn, err := db.AbreConexao(confi.DB)
    if err != nil {
      panic(err)
    }

	err = db.CriaTabela(conn)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

 	handlers.RegistrarRotas(router, conn)

	err = http.ListenAndServe(":8080", router) 
	if err != nil {
		panic(err)
	}

	
	
}
