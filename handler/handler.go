package handler

import (
	"fmt"
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
	return
}

func DeleteFile(c *gin.Context) {
	filename := c.Param("name")
	if filename == "" {
		c.String(400, "target name is empty")
		return
	}

	err := os.Remove(filepath.Join(BASEPATH, filename))
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "done")
	return
}

func UploadFile(c *gin.Context) {
	c.String(200, "upload file")
	return
}

func RenameFile(c *gin.Context) {
	newname := c.PostForm("newname")
	oldname := c.PostForm("oldname")
	if newname == "" || oldname == "" {
		c.String(400, "new(old)name is empty")
		return
	}

	fmt.Printf("newname: %s\noldname: %s\n", newname, oldname)

	err := os.Rename(filepath.Join(BASEPATH, oldname), filepath.Join(BASEPATH, newname))
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "done")
	return
}

func GetHome(c *gin.Context) {
	resp, err := getfolder(BASEPATH)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, resp)
	return
}

func GetFolder(c *gin.Context) {
	reqPath := c.Param("name")
	fullPath := filepath.Join(BASEPATH, reqPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.String(400, fmt.Sprintf("filepath %s not exist", fullPath))
		return
	}

	resp, err := getfolder(fullPath)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, resp)
	return
}

func MkFolder(c *gin.Context) {
	fn := c.Param("name")
	if fn == "" {
		c.String(400, "no new folder name")
		return
	}

	nowpath := c.PostForm("path")
	if nowpath == "" {
		c.String(400, "no now path")
		return
	}

	fullpath := filepath.Join(BASEPATH, nowpath, fn)
	err := os.Mkdir(fullpath, 0770)
	if err != nil {
		c.String(500, fmt.Sprintf("create folder: < %s > err => %s", fullpath, err))
		return
	}

	c.String(200, "done")
	return
}

func RenameFolder(c *gin.Context) {
	oldname := c.Param("oldname")
	newname := c.Param("newname")
	if oldname == "" || newname == "" {
		c.String(400, "empty old(new)name")
		return
	}

	err := os.Rename(oldname, newname)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "done")
	return
}
