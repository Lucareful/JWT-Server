package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/luenci/oauth2/service"
)

const BEARER_SCHEMA = "Bearer"

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, _ := c.Cookie("Token")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		srv := service.NewALLService()
		var token, err = srv.JWT.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
