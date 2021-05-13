package model

import (
	"giligili/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Video struct {
	gorm.Model
	Title string
	Info string
}

func (v *Video) Clicks() uint64 {
	//若视频还未被点击过，经过一系列中间转换最后返回的count为0
	countStr,_ := cache.RedisClient.Get(cache.VideoClickKey(v.ID)).Result()
	count,_ := strconv.ParseUint(countStr,10,64)
	return count
}

func (v *Video) ClicksAdd()  {
	cache.RedisClient.Incr(cache.VideoClickKey(v.ID))
	cache.RedisClient.ZIncrBy(cache.DailyRankKey,1,strconv.Itoa(int(v.ID)))
}
