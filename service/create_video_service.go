package service

import (
	"giligili/model"
	"giligili/serializer"
)

type CreateVideoService struct {
	Title string`form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string`form:"info" json:"info" binding:"min=0,max=200"`
}

func (service CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info: service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil{
		return serializer.Response{
			Code: 50001,
			Msg: "视频上传失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
		Error: "",
	}
}
