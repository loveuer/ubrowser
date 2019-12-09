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

	gFiles := ginter.Group("/file")
	{
		gFiles.GET("/:name", handler.GetFile)
	}
	gFolder := ginter.Group("/folder")
	{
		gFolder.GET("/:name", handler.GetFolder)
		gFolder.POST("/:name", handler.MkFolder)
	}

	ginter.Run(":9119")
}
