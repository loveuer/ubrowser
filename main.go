package main

import (
	"ubrowser/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	ginter := gin.Default()

	ginter.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// 操作文件
	gFiles := ginter.Group("/file")
	{
		gFiles.GET("/:name", handler.GetFile)
		gFiles.DELETE("/:name", handler.DeleteFile)
	}
	// 操作文件夹
	gFolder := ginter.Group("/folder")
	{
		gFolder.GET("/:name", handler.GetFolder)
		gFolder.POST("/:name", handler.MkFolder)
	}

	ginter.Run(":9119")
}
