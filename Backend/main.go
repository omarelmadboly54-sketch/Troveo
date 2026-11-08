package main

import (
	"Trovero/Backend/configuration"
	"Trovero/Backend/internal/controllers"
	"Trovero/Backend/internal/providers/postgres"
	"Trovero/Backend/internal/service"
	"Trovero/Backend/tokens"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db, config, err := configuration.InitPostgres()
	if err != nil {
		log.Fatal("DataBase connection Failed")
	}
	log.Println("DataBase Connection Successfully")
	tokenMaker,err:=tokens.NewJwtMaker(config.SecretKey)
	if err!=nil{
		log.Fatalf("failed to generate token maker:%v",err)
	}

	r := gin.Default()

	//User
	tokenDuration := time.Duration(config.AccessTokenMinutes)*time.Minute

	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository,tokenMaker,tokenDuration)
	userController := controllers.NewUserController(userService,tokenMaker)
	userController.RegisterRoutes(r.Group("/api"))

	if err := r.Run(config.ServerAddress); err != nil {
		log.Fatalf("failed to load serve:%v",err)
	}
}

