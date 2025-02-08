package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/timer/timer"
)

func TimerHandler(c *gin.Context) {
	response := &timer.Response{
		Hostname: os.Getenv("HOSTNAME"),
		Time:     time.Now(),
	}

	c.IndentedJSON(http.StatusAccepted, response)
}
