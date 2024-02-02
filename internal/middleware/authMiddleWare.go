package middleware

import (
	"backend/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("authHeader", authHeader)
		_, _, _, authorities, err := jwt.AccessTokenVerifier(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "UnAuthorized"})
			return
		}
		var flag bool
	Loop1:
		for _, authority := range authorities {
			for _, role := range roles {
				if role == authority.Role {
					flag = true
					break Loop1
				}
			}
		}
		if !flag {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "UnAuthorized"})
			return
		}
	}
}
