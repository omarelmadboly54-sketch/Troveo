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
