package main

import (
	"log"

	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, CoolStreamMovieServer!")
	})

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server", err)
	}
}
