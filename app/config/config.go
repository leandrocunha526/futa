package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DbHost       string `mapstructure:"DBHOST"`
	DbName       string `mapstructure:"DBNAME"`
	DbUser       string `mapstructure:"DBUSER"`
	DbPassword   string `mapstructure:"DBPASSWORD"`
	DbPort       string `mapstructure:"DBPORT"`
	JwtSecretKey string `mapstucture:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	mode := os.Getenv("APP_ENV")

	viper.AddConfigPath(".")
	if mode == "prod" {
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile("env-dev")
	}
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
