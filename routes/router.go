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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "http server ok!",
		})
	})

	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)

	file := r.Group("/file", middleware.JWTAuthMiddleware())
	{
		file.POST("/upload", controller.FileUpLoadHandler)
		file.POST("/fastupload", controller.FastFileUpLoadHandler)
		chunk := file.Group("/chunk")
		{
			chunk.POST("/init", controller.ChunkInitHandler)
			chunk.POST("/upload", controller.ChunkUpLoadHandler)
			chunk.POST("/complete", controller.ChunkCompleteHandler)
		}
		//file.POST("/delete", controller.FileDeleteHandler)
		//file.POST("/update", controller.FileUpDateHandler)
		//file.GET("/query", controller.FileQueryHandler)
		//file.GET("/download", controller.FileDownloadHandler)
	}

	return r
}
