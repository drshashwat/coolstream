package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/drshashwat/coolstream/server/CoolStreamMovieServer/controllers"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate())
}
