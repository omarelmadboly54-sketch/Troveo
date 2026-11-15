package controllers

import (
	"Trovero/Backend/internal/apperrors"
	"Trovero/Backend/internal/middleware"
	"Trovero/Backend/internal/presenters"
	"Trovero/Backend/internal/service"
	"Trovero/Backend/tokens"
	"errors"
	"net/http"

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
