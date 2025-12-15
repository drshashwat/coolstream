package main

import (
	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/routes"

	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger

func init() {
	log = logger.GetLogger()
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn().Err(err).Msg("failed to load .env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, CoolStreamMovieServer!")
	})

	routes.SetupUnprotectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
