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

type categoryController struct {
	categoryService service.CategoryService
	tokenMaker      tokens.Maker
}

type CategoryController interface {
	RegisterRoutes(g *gin.RouterGroup)
}

func NewCategoryController(categoryService service.CategoryService, tokenMaker tokens.Maker) CategoryController {
	return &categoryController{
		categoryService: categoryService,
		tokenMaker:      tokenMaker,
	}
}

func (ctl *categoryController) RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/categories", middleware.AuthMiddleware(ctl.tokenMaker), middleware.RestrictTo("admin"), ctl.CreateCategory)
	g.GET("/categories",ctl.GetAll)
	g.DELETE("/categories/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.DeleteCategoryByID)

}

func (ctl *categoryController) CreateCategory(g *gin.Context) {
	req := presenters.CreateCategoryReq{}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toCategory := req.ToCategory()

	err := ctl.categoryService.CreateCategory(&toCategory)
	if err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusCreated,gin.H{
		"message":"Created Successfully",
	})
}

func (ctl *categoryController)GetAll(g *gin.Context){
	categories,err:=ctl.categoryService.GetAll()
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{
		"categories":categories,
	})
}

func (ctl *categoryController)DeleteCategoryByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	err=ctl.categoryService.DeleteCategoryByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrCategoryNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrCategoryNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{
		"message":"Deleted successfully",
	})
}
