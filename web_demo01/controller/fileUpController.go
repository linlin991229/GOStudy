package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileUpController struct{}

func (f FileUpController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println(file.Filename)

	dst := "./" + file.Filename

	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"status": "upload success",
	})
}

func (f FileUpController) MultiUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	files := form.File["files"]

	for _, file := range files {
		log.Println(file.Filename)
		// 保存文件
		c.SaveUploadedFile(file, "./"+file.Filename)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "multi upload success",
	})
}
