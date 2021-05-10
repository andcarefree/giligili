package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UpdateVideoService struct {
	Title string`form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string`form:"info" json:"info" binding:"min=0,max=200"`
}

func (service UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err!=nil{
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg: "视频更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
