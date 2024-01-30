package db

import (
	"database/sql"
	"fmt"
	"godev/configs"

	_ "github.com/lib/pq"
)
var config  configs.DBConfig

func AbreConexao(configs.DBConfig) (*sql.DB, error){

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Porta, config.Usuario, config.Senha, config.BancoDados )

	c, err := sql.Open("postgres", sc)
	if err != nil {
		fmt.Printf("Erro ao abrir conexão")
	} 



	return c, nil
}