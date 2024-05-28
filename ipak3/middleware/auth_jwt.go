package middleware

import (
	"errors"
	"log"
	"net/http"
	"simpel-wasnaker/ipak3/helper"
	"simpel-wasnaker/ipak3/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(service service.JWTServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.ResponseFormatter(http.StatusBadRequest, "error", "Authorization needed.", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		_, err := service.ValidateToken(authHeader)
		if err != nil {
			log.Println(err)
			response := helper.ResponseFormatter(http.StatusUnauthorized, "error", errors.New("token is not valid"), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		c.Next()
	}
}

func AuthorizeUserJWT(service service.JWTServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := service.ValidateToken(authHeader)
		if err != nil {
			log.Println(err)
			response := helper.ResponseFormatter(http.StatusUnauthorized, "error", errors.New("token is not valid"), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("user", claims["user"])
		}
		c.Next()
	}
}
