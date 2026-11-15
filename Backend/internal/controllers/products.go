package controllers

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/middleware"
	"Trovero/Backend/internal/presenters"
	"Trovero/Backend/internal/service"
	"Trovero/Backend/tokens"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
	tokenMaker     tokens.Maker
}

type ProductController interface {
	RegisterRoutes(g *gin.RouterGroup)
}

func NewProductController(productService service.ProductService, tokenMaker tokens.Maker) ProductController {
	return &productController{
		productService: productService,
		tokenMaker:     tokenMaker,
	}
}

func (ctl *productController) RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/products", middleware.AuthMiddleware(ctl.tokenMaker), middleware.RestrictAccountType("admin", "seller"), ctl.CreateProduct)
	g.GET("/products",ctl.GetAll)
	g.GET("/products/:id",ctl.GetProductByID)
	g.GET("/products/mine",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictAccountType("admin","seller"),ctl.GetMyProducts)
	g.GET("/products/:id/mine",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictAccountType("admin","seller"),ctl.GetMyProduct)
	g.DELETE("/products/:id/mine",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictAccountType("admin","seller"),ctl.DeleteProductByID)
	g.GET("categories/:id/products",ctl.GetProductsByCategory)
}

func (ctl *productController) CreateProduct(g *gin.Context) {

	payload, ok := middleware.GetTokenPayload(g)
	if !ok {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request"})
		return
	}

	userID := payload.UserID

	req := presenters.CreateProductRequest{}

	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toProduct := req.ToProduct()

	fileHeader, err := g.FormFile("photo")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "product photo is required"})
		return
	}

	err = ctl.productService.CreateProduct(&toProduct, fileHeader, userID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidSize):
			g.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidSize.Error()})
		case errors.Is(err, apperrors.ErrInvalidImageType):
			g.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidImageType.Error()})
		default:
			g.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "Created Succesfully",
	})

}


func (ctl *productController)GetAll(g *gin.Context){
	page,err:=strconv.Atoi(g.DefaultQuery("page","1"))
	if err!=nil || page<1{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid page"})
		return
	}

	limit,err:=strconv.Atoi(g.DefaultQuery("limit","10"))
	if err!=nil || limit<1{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid limit"})
		return
	}

	products,err:=ctl.productService.GetAll(page,limit)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"products":products,
	})
}


func (ctl *productController)GetProductByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	product,err:=ctl.productService.GetProductByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrProductNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrProductNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
		return
	}
	g.JSON(http.StatusOK,gin.H{
		"product":product,
	})

}

func (ctl *productController)GetMyProducts(g *gin.Context){
	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}

	userID:=payload.UserID

	page,err:=strconv.Atoi(g.DefaultQuery("page","1"))
	if err!=nil || page<1{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid page"})
	}

	limit,err:=strconv.Atoi(g.DefaultQuery("limit","10"))
	if err!=nil || limit<1{
		g.JSON(http.StatusBadRequest,gin.H{"error":"Invalid limit"})
	}

	products,err:=ctl.productService.GetMyProducts(userID,page,limit)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":"internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"products":products,
	})

}


func (ctl *productController)GetMyProduct(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	paylaod,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}

	userID:=paylaod.UserID

	product,err:=ctl.productService.GetMyProduct(id,userID)
	if err!=nil{
		if errors.Is(err,apperrors.ErrProductNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrProductNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"product":product,
	})
}

func (ctl *productController)DeleteProductByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}
	userID:=payload.UserID

	err=ctl.productService.DeleteProductByID(id,userID)
	if err!=nil{
		if errors.Is(err,apperrors.ErrProductNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrProductNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}
	g.JSON(http.StatusOK,gin.H{
		"message":"Deleted successfully",
	})
}

func (ctl *productController)GetProductsByCategory(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	
	products,err:=ctl.productService.GetProductsByCategory(id)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}
	g.JSON(http.StatusOK,gin.H{
		"products":products,
	})
}
