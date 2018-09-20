package models

import (
	"time"
)

// 登录用户
type User struct {
	//主键
	Id			int	`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`
	//登录名
	Username	string	`gorm:"column:USERNAME;unique" json:"USERNAME"`
	//密码
	Password	string	`gorm:"column:PASSWORD" json:"PASSWORD"`
	//盐
	Salt		string	`gorm:"column:SALT" json:"SALT"`
	//昵称
	Nickname	string	`gorm:"column:NICKNAME" json"NICKNAME"`
}

// 文章分类
type Category struct {
	Id		int		`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`						//主键
	Title	string		`gorm:"column:TITLE;unique" json:"TITLE"`					//分类名称
	Created	time.Time	`gorm:"column:CREATED" json:"CREATED"`			//创建时间
	TopicTime time.Time	`gorm:"column:TOPIC_TIME" json:"TOPIC_TIME"`		//更新时间
	TopicCount int	`gorm:"column:TOPIC_COUNT" json:"TOPIC_COUNT"`	//分类下文章数量
	Topics	[]Topic
}

// 发表文章
type Topic struct {
	Id			int		`gorm:"column:ID;AUTO_INCREMENT;primary_key" json:"ID"`			//主键
	Uid			int		`gorm:"column:USER_ID" json:"USER_ID"`//用户id
	Title		string		`gorm:"column:TITLE" json:"TITLE"`		//文章标题
	Content		string		`gorm:"column:CONTENT;type:text" json:"CONTENT"`//文章内容
	Attachmemt	string		`gorm:"column:ATTACHMENT" json:"ATTACHMENT"`//标签
	Created		time.Time	`gorm:"column:CREATED" json"CREATED"`	//创建时间
	Updated		time.Time	`gorm:"column:UPDATED" json"UPDATED"`	//更新时间
	Views		int		`gorm:"column:VIEWS" json"VIEWS"`		//浏览数
	Author		string		`gorm:"column:AUTHOR" json"AUTHOR"`	//作者
	ReplayTime	time.Time	`gorm:"column:REPLAY_TIME" json"REPLAY_TIME"`	//最后评论时间
	ReplayCount	int		`gorm:"column:REPLAY_COUNT" json"REPLAY_COUNT"` 	//评论数
	ReplayLastUserId int	`gorm:"column:REPLAY_LAST_USER_ID" json"REPLAY_LAST_USER_ID"`	//最后评论用户Id
	CategoryId	int		`gorm:"column:CATEGORY_ID;index" json:"CATEGORY_ID"`	//所属分类
	Category 	Category
}

// 文章概述
type TopicSummary struct {
	// 主键与topic主键相同
	Id			int		`gorm:"column:ID" json:"ID"`
	// 作者
	Author		string		`gorm:"column:AUTHOR" json"AUTHOR"`
	// 概述内容
	Content		string		`gorm:"column:CONTENT" json:"CONTENT"`
	//
	Title		string		`gorm:"column:TITLE" json:"TITLE"`
	//// 评论数
	//CommentCount int		`gorm:"column:COMMENT_COUNT" json:"COMMENT_COUNT"`
	//// 浏览数
	//Views		int		`gorm:"column:VIEWS" json:"VIEWS"`
	// 创建日期
	Created		time.Time	`gorm:"column:CREATED" json:"CREATED"`
}

// 事件实体
type EventEntity struct {
	// 主键
	Id			int		`gorm:"column:ID;AUTO_INCREMENT;primary_key"		json:"ID"`
	// 被评论或点赞主体类型：默认为文章0
	EntityType	int8		`gorm:"column:ENTITY_TYPE" json:"ENTITY_TYPE"`
	// 被评论或点赞主题Id
	EntityId	int		`gorm:"column:ENTITY_ID" json:"ENTITY_ID"`
	// 评论或点赞人ID
	ActorId		int		`gorm:"column:ACTOR_ID" json:"ACTOR_ID"`
	// 被评论或点赞主体所属人
	ReceiverID	int		`gorm:"column:RECEIVER_ID" json:"RECEIVER_ID"`
	// 评论内容，如果是点赞则为空
	Content		string		`gorm:"column:CONTENT" json:"CONTENT"`
	// 评论或点赞时间
	Created		time.Time	`gorm:"column:CREATED" json:"CREATED"`
	// 父评论ID 如果没有则为0，如果是点赞则为-1
	ParentId	int		`gorm:"column:PARENT_ID" json:"PARENT_ID"`
}

// 评论
type Comment struct {
	// 主键
	Id			int		`gorm:"column:ID;AUTO_INCREMENT;primary_key"  json:"ID"`
	// 文章Id
	TopicId		int		`gorm:"column:TOPICID"		json:"TOPICID"`
	// 评论内容
	Content		string		`gorm:"column:CONTENT"		json:"CONTENT"`
	// 评论人邮箱
	Email		string		`gorm:"column:EMAIL"			json:"EMAIL"`
	// 评论人
	Name		string		`gorm:"column:NAME"			json:"NAME"`
	// 父评论，如果没有则为0
	ParentId	int		`gorm:"column:PARENT_ID"	json:"PARENT_ID"`
	// 创建时间
	Created		time.Time	`gorm:"column:CREATED"		json:"CREATED"`
}