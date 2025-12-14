package routes

import (
	controller "github.com/drshashwat/coolstream/server/CoolStreamMovieServer/controllers"
	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
}
