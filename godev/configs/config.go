package configs

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

var conf *config

type config struct {
	API APIConfig
	DB DBConfig
}

type APIConfig struct {
	Porta string
}

type DBConfig struct {
	Host string
	Porta string
	Usuario string
	Senha string
	BancoDados string

}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Carrega() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	conf = new(config)

	conf.API = APIConfig{
		Porta: viper.GetString("api.port"),

	}

	conf.DB = DBConfig{
		Host: viper.GetString("database.host"),
		Porta: viper.GetString("database.port"),
		Usuario: viper.GetString("database.usuario"),
		Senha: viper.GetString("database.senha"),
		BancoDados: viper.GetString("database.nome"),
	}

	return nil
}

func GetDB() DBConfig {
	return conf.DB
}

func GetServerPort() string {
	return conf.API.Porta
}


func RespondWithError(w http.ResponseWriter, code, errorCode int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "code": errorCode})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}



