package service

import "gitlab.com/clickertech/gogin/src/model"

type VideoService interface {
	FindAll() []model.Video
	Save(model.Video) model.Video
}

type videoService struct {
	videos []model.Video
}

// Constrctor
func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video model.Video) model.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []model.Video {
	return service.videos
}
