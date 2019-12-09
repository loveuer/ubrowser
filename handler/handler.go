package handler

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	BASEPATH string
	err      error
)

func init() {
	BASEPATH, err = os.Getwd()
	if err != nil {
		log.Fatalf("< handler > < init > os get wd err => %v\n", err)
	}
	BASEPATH = filepath.Join(BASEPATH, "Files")
}

func GetFile(c *gin.Context) {
	c.File(filepath.Join(BASEPATH, c.Param("name")))
}

func GetFolder(c *gin.Context) {
	reqPath := c.Param("name")
	fullPath := filepath.Join(BASEPATH, reqPath)

	resp, err := getfolder(fullPath)
	if err != nil {
		c.String(500, err.Error())
	}

	c.JSON(200, resp)
}

func UploadFile(c *gin.Context) {
	c.String(200, "upload file")
}
