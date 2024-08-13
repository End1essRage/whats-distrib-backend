package main

import (
	"fmt"
	"net/http"

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
}

func (a *Api) uploadFile(c *gin.Context) {
	logrus.Info("hitted upload file")

	//taking
	file, _ := c.FormFile("file")

	//saving
	if err := c.SaveUploadedFile(file, "uploaded/"+file.Filename); err != nil {
		logrus.Error("error while saving file : " + err.Error())
	}

	//smth
	a.service.HandleScanRequest(file.Filename)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
