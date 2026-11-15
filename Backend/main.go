package main

import (
	"Trovero/Backend/configuration"
	"Trovero/Backend/internal/controllers"
	"Trovero/Backend/internal/providers/postgres"
	"Trovero/Backend/internal/service"
	"Trovero/Backend/tokens"
	"log"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, config, err := configuration.InitPostgres()
	if err != nil {
		log.Fatal("DataBase connection Failed")
	}
	log.Println("DataBase Connection Successfully")

	tokenMaker, err := tokens.NewJwtMaker(config.SecretKey)
	if err != nil {
		log.Fatalf("failed to generate token maker:%v", err)
	}

	cld, err := cloudinary.NewFromURL(config.CloudinaryUrl)
	if err != nil {
		log.Fatalf("failed to initialize cloudinary:%v",err)
	}
	log.Println("Cloudinary Initialized Successfully")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.FrontendUrl},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
	}))

	//User
	tokenDuration := time.Duration(config.AccessTokenMinutes) * time.Minute

	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository, tokenMaker, tokenDuration)
	userController := controllers.NewUserController(userService, tokenMaker)
	userController.RegisterRoutes(r.Group("/api"))
	//Address
	addressRepository := postgres.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepository, userService)
	addressController := controllers.NewAddressController(addressService, tokenMaker)
	addressController.RegisterRoutes(r.Group("/api"))

	//Category
	categoryRepository := postgres.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService, tokenMaker)
	categoryController.RegisterRoutes(r.Group("/api"))

	//Product
	productRepository:=postgres.NewProductRepository(db)
	productService:=service.NewProductService(productRepository,cld)
	productController:=controllers.NewProductController(productService,tokenMaker)
	productController.RegisterRoutes(r.Group("/api"))
	//Coupon
	couponRepository:=postgres.NewCouponRepository(db)
	couponService:=service.NewCouponService(couponRepository)
	couponController:=controllers.NewCouponController(couponService,tokenMaker)
	couponController.RegisterRoutes(r.Group("/api"))

	if err := r.Run(config.ServerAddress); err != nil {
		log.Fatalf("failed to load serve:%v", err)
	}
}
