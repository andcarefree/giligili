package serializer

import (
	"giligili/model"
)

//Video 视频序列化器
type Video struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Info string `json:"info"`
	CreatAt int64 `json:"creat_at"`
	Clicks uint64
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID: item.ID,
		Title: item.Title,
		Info: item.Info,
		CreatAt: item.CreatedAt.Unix(),
		Clicks: item.Clicks(),
	}
}

func BuildVideos(items []model.Video) (videos []Video) {
	for _,item :=range items {
		video := BuildVideo(item)
		videos = append(videos,video)
	}
	return
}
