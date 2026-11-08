package utils

import (
	"Trovero/Backend/internal/models"

	"github.com/spf13/viper"
)

func LoadConfiguration(path string) (config models.Configuration, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err=viper.ReadInConfig()
	if err!=nil{
		return 
	}

	err=viper.Unmarshal(&config)
	return
	
}
