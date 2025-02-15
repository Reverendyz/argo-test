package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reverendyz/timer/handlers"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"*"},
	}))
	router.GET("/", handlers.Heathz)
	router.GET("/timer", handlers.TimerHandler)

	router.Run(fmt.Sprintf("%s:%s", host, port))
}
