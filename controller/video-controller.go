package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/clickertech/gogin/src/model"
	"gitlab.com/clickertech/gogin/src/service"
)

type VideoController interface {
	FindAll() []model.Video
	Save(c *gin.Context) model.Video
}

type controller struct {
	service service.VideoService
}

// Constrctor
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) model.Video {
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
