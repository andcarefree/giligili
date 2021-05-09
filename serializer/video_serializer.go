package serializer

import "giligili/model"

//Video 视频序列化器
type Video struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Info string `json:"info"`
	CreatAt int64 `json:"creat_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID: item.ID,
		Title: item.Title,
		Info: item.Info,
		CreatAt: item.CreatedAt.Unix(),
	}
}
