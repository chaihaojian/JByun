package routes

import (
	"JByun/controller"
	"JByun/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	register := r.Group("/register")
	{
		register.POST("/submit", controller.SubmitHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "http server ok!",
		})
	})

	return r
}
