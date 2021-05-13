package service

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"strings"
)

type DailyRankService struct {

}

func (service *DailyRankService)Get() serializer.Response {
	var videos []model.Video

	//从redis的有序集合中读取点击量前十的视频
	vids,_ := cache.RedisClient.ZRevRange(cache.DailyRankKey,0,9).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id,%s)",strings.Join(vids,","))
		err := model.DB.Where("id in (?)",vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code: 50000,
				Msg: "数据库查询错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
