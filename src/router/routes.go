package router

import (
	"booleanservice/src/controller"
	"booleanservice/src/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//SetupRouter is
func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware, err := service.JWTMiddleware()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	auth := r
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/", controller.CreateValue)
		auth.DELETE("/:id", controller.DeleteValue)
		auth.PATCH("/:id", controller.UpdateValue)
		auth.GET("/:id", controller.GetValue)
	}
	return auth
}
