package serializer

import (
	"giligili/model"
)

//Video 视频序列化器
type Video struct {
	ID uint `json:"id"`
	AuthNickName string
	AuthName string
	Title string `json:"title"`
	Info string `json:"info"`
	CreatAt int64 `json:"creat_at"`
	VideoURL string `json:"video_url"`
	PosterURL string `json:"poster_url"`
	Clicks uint64
}

func BuildVideo(item model.Video) Video {
	var auth model.User
	model.DB.Find(&auth,item.AuthID)
	return Video{
		ID: item.ID,
		AuthNickName: auth.Nickname,
		AuthName: auth.UserName,
		Title: item.Title,
		Info: item.Info,
		CreatAt: item.CreatedAt.Unix(),
		Clicks: item.Clicks(),
		VideoURL: item.VideoURL,
		PosterURL: item.PosterURL,
	}
}

func BuildVideos(items []model.Video) (videos []Video) {
	for _,item :=range items {
		video := BuildVideo(item)
		videos = append(videos,video)
	}
	return
}
