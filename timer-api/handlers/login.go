package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/reverendyz/timer/utils"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtCustomToken struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	var payload LoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	if payload.Username == "admin" && payload.Password == "password" {
		now := time.Now()
		claims := JwtCustomToken{
			Username: payload.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "revzwebapp.com",
				Subject:   "login",
				Audience:  []string{"developers", "students", "hackers"},
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(1 * time.Hour)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
		tokenString, err := token.SignedString(utils.JwtPrivateKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar o token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
}
