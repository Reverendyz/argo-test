package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomToken struct {
	Email string `json:"email"`
	Roles jwt.ClaimStrings
	jwt.RegisteredClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &JwtCustomToken{}, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodEdDSA {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return JwtPublicKey, nil
		})
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(*JwtCustomToken)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
			return
		}

		now := time.Now()

		if claims.ExpiresAt != nil && now.After(claims.ExpiresAt.Time) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
			return
		}

		if claims.NotBefore != nil && now.Before(claims.NotBefore.Time) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token haven't been validated"})
			return
		}

		if claims.Issuer != "revzwebapp.com" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Issuer"})
			return
		}

		if claims.Subject != "login" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Subject"})
			return
		}

		validAudience := false
		expectedAudiences := []string{"developers", "students", "hackers"}
		for _, expectedAud := range expectedAudiences {
			for _, aud := range claims.Audience {
				if aud == expectedAud {
					validAudience = true
					break
				}
			}
		}
		if !validAudience {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Audience"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
