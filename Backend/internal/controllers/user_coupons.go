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

type usercouponController struct {
	usercouponService service.UserCouponService
	tokenMaker        tokens.Maker
}

type UserCouponController interface {
	RegisterRoutes(g *gin.RouterGroup)
}

func NewUserCouponController(usercouponService service.UserCouponService, tokenMaker tokens.Maker) UserCouponController {
	return &usercouponController{
		usercouponService: usercouponService,
		tokenMaker:        tokenMaker,
	}
}

func (ctl *usercouponController) RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/coupons/users", middleware.AuthMiddleware(ctl.tokenMaker), middleware.RestrictTo("admin"), ctl.AssignCouponToUser)
	g.GET("/coupons/mine", middleware.AuthMiddleware(ctl.tokenMaker), middleware.RestrictAccountType("customer"), ctl.GetMyCoupons)
	g.GET("/coupons/mine/:id",middleware.AuthMiddleware(ctl.tokenMaker),middleware.RestrictAccountType("customer"),ctl.GetMyCoupon)
}

func (ctl *usercouponController) AssignCouponToUser(g *gin.Context) {

	req := presenters.AssignCouponToUserReq{}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	toUserCoupon := req.ToUserCoupon()

	err := ctl.usercouponService.AssignCouponToUser(&toUserCoupon)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrUserNotFound):
			g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		case errors.Is(err, apperrors.ErrCouponNotFound):
			g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		case errors.Is(err, apperrors.ErrCouponAssignmentToNonCustomer),
			errors.Is(err, apperrors.ErrCouponNotActive),
			errors.Is(err, apperrors.ErrCouponExpired):
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		case errors.Is(err, apperrors.ErrCouponAlreadyAssigned):
			g.JSON(http.StatusConflict, gin.H{"error": err.Error()})

		default:
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})

		}
		return
	}

	g.JSON(http.StatusCreated, gin.H{
		"message": "Coupon Assigned Successfully",
	})

}

func (ctl *usercouponController) GetMyCoupons(g *gin.Context) {

	page, err := strconv.Atoi(g.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	limit, err := strconv.Atoi(g.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	paylad, ok := middleware.GetTokenPayload(g)
	if !ok {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request"})
		return
	}

	userID := paylad.UserID

	userCoupons, err := ctl.usercouponService.GetMyCoupons(userID, page, limit)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"userCoupons": userCoupons,
	})

}


func (ctl *usercouponController)GetMyCoupon(g *gin.Context){
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

	userCoupon,err:=ctl.usercouponService.GetMyCoupon(id,userID)
	if err!=nil{
		if errors.Is(err,apperrors.ErrCouponNotFound){
			g.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError,gin.H{"error":"Internal server error"})
		return
	}

	g.JSON(http.StatusOK,gin.H{
		"userCoupon":userCoupon,
	})
	
}
