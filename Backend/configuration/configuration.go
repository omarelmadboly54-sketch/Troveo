package configuration

import (
	"Trovero/Backend/internal/models"
	"Trovero/Backend/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB,models.Configuration,error) {
	config,err:=utils.LoadConfiguration("./Backend")
	if err!=nil{
		return nil,models.Configuration{},err
	}

	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})
	if err!=nil{
		return nil,models.Configuration{},err
	}
	
	return db,config,nil

}

