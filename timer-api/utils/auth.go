package utils

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	validTokens = make(map[string]struct{})
	tokenLock   = sync.RWMutex{}
)

func AddValidToken(token string) {
	tokenLock.Lock()
	defer tokenLock.Unlock()
	validTokens[token] = struct{}{}
}

func IsValidToken(token string) bool {
	tokenLock.RLock()
	defer tokenLock.RUnlock()
	_, exists := validTokens[token]
	return exists
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if !IsValidToken(authHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
