package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Api struct {
	service Service
	gin     *gin.Engine
}

func NewApi(s Service) *Api {
	api := &Api{gin: gin.Default(), service: s}
	api.gin.MaxMultipartMemory = 8 << 20 // 8MiB
	api.initRoutes()

	return api
}

func (a *Api) initRoutes() {
	a.gin.GET("/status", func(c *gin.Context) {
		logrus.Info("hitted /status")
		c.JSON(http.StatusAccepted, nil)
	})

	a.gin.POST("/upload", a.uploadFile)
	a.gin.POST("/clear", a.clearUploaded)
}

func (a *Api) uploadFile(c *gin.Context) {
	logrus.Info("hitted upload file")

	//taking
	file, _ := c.FormFile("file")

	name := GenerateFileName(file.Filename)
	//saving
	if err := c.SaveUploadedFile(file, "uploaded/"+name); err != nil {
		logrus.Error("error while saving file : " + err.Error())
	}

	//smth
	a.service.HandleScanRequest(name)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (a *Api) clearUploaded(c *gin.Context) {
	logrus.Info("hitted clear uploaded")

	if err := os.RemoveAll("uploaded"); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, "removed")
}
