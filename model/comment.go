package model

import "github.com/jinzhu/gorm"

//评论总体来说分为两种
//一种是顶层评论，这种回复只需要通知视频作者
//二种是楼中楼回复，通知作者,同时递归通知回复评论（楼中楼回复即为回复顶层评论）

//TODO @功能应该单独抽象出一个函数

type Comment struct {
	gorm.Model
	//表示评论是否为顶层评论
	IsTop bool
	//评论关联的video
	VideoID uint
	//评论的用户
	UserID uint
	//回复的评论ID
	ReplyID uint
	//评论的正文
	content string
}