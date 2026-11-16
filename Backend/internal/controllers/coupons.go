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

type couponController struct {
	couponService service.CouponService
	tokenMaker    tokens.Maker
}

type CouponController interface {
	RegisterRoutes(g *gin.RouterGroup)
}

func NewCouponController(couponService service.CouponService, tokenMaker tokens.Maker) CouponController {
	return &couponController{
		couponService: couponService,
		tokenMaker:    tokenMaker,
	}
}

func (ctl *couponController) RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/coupons", middleware.AuthMiddleware(ctl.tokenMaker), middleware.RestrictTo("admin"), ctl.CreateCoupon)
	g.GET("/coupons",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.GetAll)
	g.GET("/coupons/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.GetCouponByID)
	g.DELETE("/coupons/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.DeleteCouponByID)
	g.PATCH("/coupons/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictTo("admin"),ctl.UpdateCouponStatus)
	
}

func (ctl *couponController) CreateCoupon(g *gin.Context) {
	req := presenters.CreateCouponReq{}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toCoupon := req.ToCoupon()

	err := ctl.couponService.CreateCoupon(&toCoupon)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidCouponValue):
			g.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidCouponValue.Error()})
		case errors.Is(err, apperrors.ErrInvalidCouponExpiration):
			g.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidCouponExpiration.Error()})
		case errors.Is(err, apperrors.ErrCouponCodeExists):
			g.JSON(http.StatusConflict, gin.H{"error": apperrors.ErrCouponCodeExists.Error()})
		default:
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server error"})
		}
		return
	}
	g.JSON(http.StatusCreated, gin.H{
		"message": "Created successfully",
	})
}

func (ctl *couponController)GetAll(g *gin.Context){
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

	coupons,err:=ctl.couponService.GetAll(page,limit)
	if err!=nil{
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"coupons":coupons,
	})

}

func (ctl *couponController)GetCouponByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	coupon,err:=ctl.couponService.GetCouponByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrCouponNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrCouponNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}
	g.JSON(http.StatusOK,gin.H{
		"coupon":coupon,
	})
}


func (ctl *couponController)DeleteCouponByID(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	err=ctl.couponService.DeleteCouponByID(id)
	if err!=nil{
		if errors.Is(err,apperrors.ErrCouponNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrCouponNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"message":"Deleted sucessfully",
	})
}

func (ctl *couponController)UpdateCouponStatus(g *gin.Context){
	id,err:=strconv.ParseInt(g.Param("id"),10,64)
	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	req:=presenters.UpdateCouponStatusReq{}

	if err:=g.ShouldBindJSON(&req);err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	err=ctl.couponService.UpdateCouponStatus(id,*req.IsActive)
	if err!=nil{
		if errors.Is(err,apperrors.ErrCouponNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":apperrors.ErrCouponNotFound.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"message":"Updated successfully",
	})

}