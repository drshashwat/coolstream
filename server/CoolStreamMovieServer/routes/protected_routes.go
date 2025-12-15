package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/drshashwat/coolstream/server/CoolStreamMovieServer/controllers"
	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/middleware"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.GET("/recommendedmovies", controller.GetRecomendedMovies())
}
