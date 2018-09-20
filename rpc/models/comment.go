package rpcModels

import (
	"time"
)

// 添加评论
// 评论请求表单实体
type CommentInfo struct {
	TopicId			int
	Content			string
	Name			string
	Email			string
	ParentId		int
}

type AddCommentReply struct {
	baseReply
}

// 获取评论
type GetCommentsParams struct {
	TopicId			int
	//PageNum			int
	//PageSize		int
}

type CommentReply struct {
	Id				int
	TopicId			int
	Content			string
	Name			string
	ParentId		int
	Created			time.Time
}

type GetCommentsReply struct {
	baseReply
	List			[]CommentReply
	Total			int
}

// 删除评论
type DelCommentParam struct {
	Id				int
	TopicId 		int
}

type DelCommentReply struct {
	baseReply
}

