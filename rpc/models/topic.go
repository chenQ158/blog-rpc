package rpcModels

import (
	"time"
)

// 文章提交表单实体
type TopicForm struct {
	Id				int		`form:"id"`
	Title 			string		`form:"title"`
	CatId			int		`form:"category"`
	Content			string		`form:"content"`
	Attachment		string		`form:"attachment"`
	Summary			string		`form:"summary"`
	Author			string		`form:"author"`
}

// 文章详情，用于列表展示
type TopicInfo struct {
	Id			int		`json:"ID"`
	Title		string		`json:"TITLE"`		//文章标题
	Attachmemt	string		`json:"ATTACHMENT"`//标签
	Updated		time.Time	`json:"UPDATED"`	//更新时间
	ReplayCount	int		`json:"REPLAY_COUNT"`		//评论数
	CategoryId	int		`json:"CATEGORY_ID"`
	CategoryTitle  string	`json:"CATEGORY_TITLE"`
	Uid			int		`json:"UID"`
	Author		string		`json:"AUTHOR"`
}

// 文章显示，用于首页分页显示文章
type TopicSummary struct {
	Id			int
	Title		string
	Content		string
	Created		time.Time
	Attachment	string
	Author		string
}

// 文章删除
type DelTopicParams struct {
	Id 			int
	CatId		int
}

type DelTopicReply struct {
	baseReply
}

// 文章分页获取
type GetTopicsParams struct {
	PageNum		int
	PageSize	int
}

// 文章详情信息分页
type GetTopicsReply struct {
	baseReply
	List		[]TopicInfo
	Total		int
}

// 文章简要信息分页
type GetSummarysReply struct {
	baseReply
	List		[]TopicSummary
	Total		int
}

// 指定文章获取
type GetTopicParam struct {
	Id			int
}

// 获取单篇文章详情
type GetTopicDetailsReply struct {
	baseReply
	TopicInfo
	Content		string
}

type GetTopic struct {
	baseReply
	TopicSummary
}

// 搜索文章标题关键字
type SearTopicParams struct {
	Keyword		string
}

type AddTopicReply struct {
	baseReply
}

type UpdateTopicReply struct {
	baseReply
}

type GetTopicsByCatParams struct {
	CatId		int
	PageNum		int
	PageSize	int
}