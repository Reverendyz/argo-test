package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Heathz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
