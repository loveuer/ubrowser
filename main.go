package main

import (
	"ubrowser/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	ginter := gin.Default()
	ginter.LoadHTMLFiles("templates/home.html")		

	ginter.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
		return
	})
	// 支持这个服务的网页
	ginter.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
		return
	})
	// 操作文件
	gFiles := ginter.Group("/file")
	{
		gFiles.GET("/:name", handler.GetFile)
		gFiles.DELETE("/:name", handler.DeleteFile)
		gFiles.PUT("", handler.RenameFile)
		gFiles.POST("", handler.UploadFile)
	}
	// 操作文件夹
	gFolder := ginter.Group("/folder")
	{
		gFolder.GET("", handler.GetHome)
		gFolder.GET("/:name", handler.GetFolder)
		gFolder.POST("/:name", handler.MkFolder)
	}

	ginter.Run(":9119")
}
