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

type addressController struct {
	addressService  service.AddressService
	tokenMaker		tokens.Maker
}

type AddressController interface{
	RegisterRoutes(g *gin.RouterGroup)
}

func NewAddressController (addressService service.AddressService,tokenMaker tokens.Maker)AddressController{
	return &addressController{
		addressService: addressService,
		tokenMaker: tokenMaker,
	}
}

func(ctl *addressController)RegisterRoutes(g *gin.RouterGroup){
	g.POST("/addresses",middleware.AuthMiddleware(ctl.tokenMaker),ctl.CreateAddress)
	g.GET("/addresses",middleware.AuthMiddleware(ctl.tokenMaker),ctl.GetAll)
	g.DELETE("/addresses/:id",middleware.AuthMiddleware(ctl.tokenMaker),ctl.DeleteAddressByID)
}

func(ctl *addressController)CreateAddress(g *gin.Context){

	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}
	userID:=payload.UserID

	req:=presenters.CreateAddressReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	toAddress:=req.ToAddress()
	
	err:=ctl.addressService.CreateAddress(&toAddress,userID)
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

	g.JSON(http.StatusCreated,gin.H{
		"message":"Created successfully",
	})
}


func(ctl *addressController)GetAll(g *gin.Context){

	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}

	userID:=payload.UserID

	addresses,err:=ctl.addressService.GetAll(userID)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
	})
	return
	}
	g.JSON(http.StatusOK,gin.H{
		"addresses":addresses,
	})

}

func(ctl *addressController)DeleteAddressByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error(),
	})
	return
	}

	payload,ok:=middleware.GetTokenPayload(g)
	if !ok{
		g.JSON(http.StatusUnauthorized,gin.H{"error":"unauthorized request"})
		return
	}

	userID:=payload.UserID

	err=ctl.addressService.DeleteAddressByID(id,userID)
	if err!=nil{
		if errors.Is(err,apperrors.ErrAddressNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrAddressNotFound.Error(),
		})
		return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"message":"Deleted Successfully",
	})

}
