package server

import (
	"blog-rpc/models"
	"blog-rpc/rpc/models"
	"blog-rpc/service"
	log "code.google.com/log4go"
	"context"
	"time"
)

type Comment struct {}

func (c *Comment) AddComment(ctx context.Context, req *rpcModels.CommentInfo, reply *rpcModels.AddCommentReply) error {
	log.Debug("service层<添加评论>方法调用请求:%v", req)
	var comment models.Comment
	comment.TopicId = req.TopicId
	comment.Name = req.Name
	comment.Email = req.Email
	comment.Content = req.Content
	comment.ParentId = req.ParentId
	comment.Created = time.Now()
	err := service.AddComment(&comment)
	if err != nil {
		// 根据返回的错误编码
		log.Error("service层<添加评论>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	log.Debug("添加了评论：%v", comment)
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (c *Comment) GetCommentsByTopicId(ctx context.Context, req *rpcModels.GetCommentsParams, reply *rpcModels.GetCommentsReply) error {
	log.Debug("service层<根据文章Id获取评论>方法调用请求:%v", req)
	comments, err := service.GetCommentsByTopicId(req.TopicId)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<根据文章Id获取评论>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}

	var list []rpcModels.CommentReply
	for _,comm := range *comments {
		var comReply rpcModels.CommentReply
		comReply.TopicId = comm.TopicId
		comReply.Created = comm.Created
		comReply.ParentId = comm.ParentId
		comReply.Content = comm.Content
		comReply.Name = comm.Name
		comReply.Id = comm.Id
		list = append(list, comReply)
	}
	reply.List = list
	reply.Total = len(list)
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (c *Comment) DelCommentById(ctx context.Context, req *rpcModels.DelCommentParam, reply *rpcModels.DelCommentReply) error {
	log.Debug("service层<根据文章Id获取评论>方法调用请求:%v", req)
	err := service.DelCommentById(req.Id, req.TopicId)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<根据Id删除评论>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}

	reply.ResMsg = ""
	reply.ResCode = ""
	return nil
}
