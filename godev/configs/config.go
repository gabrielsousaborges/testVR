package configs

import (
	"github.com/spf13/viper"
)



type Config struct {
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

func Carrega() (Config, error) {

	var conf *Config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return *conf, err
		}
	}

	conf = new(Config)

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

	return *conf, nil
}


