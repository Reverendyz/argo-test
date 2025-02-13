package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/timer/timer"
)

func TimerHandler(c *gin.Context) {
	t := &timer.Response{}

	response := t.GetTime()

	c.IndentedJSON(http.StatusAccepted, response)
}
