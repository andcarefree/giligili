package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ListVideosService struct {

}

func (service ListVideosService)List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err!=nil{
		return serializer.Response{
			Code: 50002,
			Msg: "数据库为空",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}