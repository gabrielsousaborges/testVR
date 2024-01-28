package db

import (
	"database/sql"
	"fmt"
	"godev/configs"

	_ "github.com/lib/pq"
)

func AbreConexao() (*sql.DB){
	config := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Porta, config.Usuario, config.Senha, config.BancoDados )

	c, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	} 



	return c
}