package routes

import (
	"JByun/controller"
	"JByun/logger"
	"JByun/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Setup 注册路由
func Setup() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "http server ok!",
		})
	})

	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	return r
}
