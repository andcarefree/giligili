package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"sort"
)

type DailyRankService struct {

}

func (service *DailyRankService)Get() serializer.Response {
	var videos []model.Video

	//从redis的有序集合中读取点击量前十的视频id
	vids,_ := cache.RedisClient.ZRevRange(cache.DailyRankKey,0,9).Result()

	//因为点击量相关数据只在redis中维护，所以返回点击量前十的视频时候需要查找数据库中的视频信息
	if len(vids) > 1 {
		//order := fmt.Sprintf("FIELD(id,%s)",strings.Join(vids,","))
		err := model.DB.Where("id in (?)",vids).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code: 50000,
				Msg: "数据库查询错误",
				Error: err.Error(),
			}
		}
	}
	//当redis中维护的点击量相同时，可以将发布时间更靠后的视频放在前面
	//这一部分排序逻辑只需要在数据库那里实现，但是我不太会
	//现在这种处理方式会导致查询redis的次数暴增
	sort.Slice(videos, func(i, j int) bool {
		if videos[i].Clicks() == videos[j].Clicks() {
			return videos[i].CreatedAt.After(videos[j].CreatedAt)
		}
		return videos[i].Clicks()>videos[j].Clicks()
	})
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
