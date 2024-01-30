package db

import (
	"database/sql"
	"fmt"
	"godev/configs"

	_ "github.com/lib/pq"
)

func AbreConexao(config configs.DBConfig) (*sql.DB, error){

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Porta, config.Usuario, config.Senha, config.BancoDados )

	c, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, err
	} 



	return c, nil
}