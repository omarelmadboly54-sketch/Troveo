package middleware

import (
	"Trovero/Backend/tokens"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker tokens.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func GetTokenPayload(g *gin.Context) (*tokens.Payload, bool) {
	val, exists := g.Get("authorization_payload")
	if !exists {
		return nil, false
	}
	payload, ok := val.(*tokens.Payload)
	return payload, ok
}


func RestrictTo(allowedRoles ...string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
       
        val, exists := ctx.Get(authorizationPayloadKey)
        if !exists {
            err := errors.New("authorization payload not found")
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }

        
        payload, ok := val.(*tokens.Payload)
        if !ok {
            err := errors.New("invalid authorization payload type")
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

   
        isAllowed := false
        for _, role := range allowedRoles {
            if payload.Role == role {
                isAllowed = true
                break
            }
        }

        if !isAllowed {
            err := errors.New("permission denied: insufficient permissions")
            ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
            return
        }

        ctx.Next()
	}
}	