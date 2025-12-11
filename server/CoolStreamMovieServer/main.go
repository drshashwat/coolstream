package main

import (
	"log"

	"github.com/gin-gonic/gin"

	controller "github.com/drshashwat/coolstream/server/CoolStreamMovieServer/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, CoolStreamMovieServer!")
	})
	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server", err)
	}
}
