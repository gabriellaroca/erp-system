package middleware

import (
	"net/http"

	authentication "erp-system.com/v1/domain/Security/Authentication/jwt"
	"erp-system.com/v1/domain/Security/Authentication/jwt/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &authentication.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return config.SecretJwt, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*authentication.CustomClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
