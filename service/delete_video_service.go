package service

import (
	"giligili/model"
	"giligili/serializer"
)

type DeleteVideoService struct {
	ID uint `form:"id" json:"id" binding:"required" parm`
}

func (service DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code: 50001,
			Msg: "视频删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}