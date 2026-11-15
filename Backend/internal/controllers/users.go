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

type userController struct {
	UserService service.UserService
	tokenMaker	tokens.Maker
}

type UserController interface{
	RegisterRoutes(g *gin.RouterGroup)
}

func NewUserController (UserService service.UserService,tokenMaker tokens.Maker)UserController{
	return &userController{
		UserService: UserService,
		tokenMaker: tokenMaker,
	}
}


func(ctl *userController)RegisterRoutes(g *gin.RouterGroup){
	g.POST("/users",ctl.CreateUser)
	g.POST("/login",ctl.Login)
	g.GET("/users/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.GetUserByID)
	g.DELETE("/users/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.DeleteUserByID)
	g.GET("/users",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.GetAll)
	g.PATCH("/users/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.UpdateUserStatus)
}

func(ctl *userController)CreateUser(g *gin.Context){
	req:=presenters.CreateUserReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	toUser:=req.ToUser()

	err:=ctl.UserService.CreateUser(&toUser)
	if err!=nil{
		if errors.Is(err,apperrors.ErrUserAlreadyExists){
			g.JSON(http.StatusBadRequest,gin.H{"error":apperrors.ErrUserAlreadyExists.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusCreated,gin.H{
		"message":"Account created successfully",
	})
}

func(ctl *userController)Login(g *gin.Context){
	req:=presenters.LoginReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	accessToken,err:=ctl.UserService.Login(req.Email,req.Password)
	if err!=nil{
		switch{
		case errors.Is(err,apperrors.ErrInvalidCredentials):
			g.JSON(http.StatusUnauthorized,gin.H{"error":apperrors.ErrInvalidCredentials.Error()})
		case errors.Is(err,apperrors.ErrUserNotAllowed):
			g.JSON(http.StatusForbidden,gin.H{"error":apperrors.ErrUserNotAllowed.Error()})
		default:
			g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})	
		}
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"accessToken":accessToken,
	})
}

func(ctl *userController)GetUserByID(g *gin.Context){

	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	user,err:=ctl.UserService.GetUserByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrUserNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrUserNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}

	g.JSON(http.StatusOK,gin.H{
		"user":user,
	})

}

func(ctl *userController)DeleteUserByID(g *gin.Context){

	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	err=ctl.UserService.DeleteUserByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrUserNotFound){
			g.JSON(http.StatusBadRequest,gin.H{"error":apperrors.ErrUserNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{
		"message":"Account deleted Successfully",
	})
}


func(ctl *userController)GetAll(g *gin.Context){

	page, err := strconv.Atoi(g.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	limit, err := strconv.Atoi(g.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	accountType := g.Query("account_type")

	users,err:=ctl.UserService.GetAll(page,limit,accountType)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{
		"users":users,
	})

}


func(ctl *userController)UpdateUserStatus(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}
	
	req:=presenters.UpdateUserStatusReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	err=ctl.UserService.UpdateUserStatus(id,*req.IsActive)
	if err!=nil{
		if errors.Is(err,apperrors.ErrUserNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrUserNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{
		"message":"Account status updated successfully",
	})

}

