package server

import (
	"blog-rpc/models"
	"blog-rpc/rpc/models"
	"blog-rpc/service"
	log "code.google.com/log4go"
	"context"
	"time"
)

type Topic struct {}

func (t *Topic) DelTopic(ctx context.Context, req *rpcModels.DelTopicParams, reply *rpcModels.DelTopicReply) error {
	err := service.DelTopic(req.Id, req.CatId)
	if err != nil {
		log.Error("service层<删除文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.ResMsg = ""
	reply.ResCode = ""
	return nil
}

func (t *Topic) GetSummarysByOffsetAndLimit(ctx context.Context, req *rpcModels.GetTopicsParams, reply *rpcModels.GetSummarysReply) error {
	infos, err := service.GetSummarysByOffsetAndLimit(req.PageNum, req.PageSize)
	if err != nil {
		log.Error("service层<分页获取文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	total := service.GetTopicsCount()
	var list []rpcModels.TopicSummary
	for _, info := range *infos {
		var sinfo rpcModels.TopicSummary
		sinfo.Id = info.Id
		sinfo.Title = info.Title
		sinfo.Author = info.Author
		sinfo.Content = info.Content
		sinfo.Created = info.Created
		list = append(list, sinfo)
	}
	reply.List = list
	reply.Total = total
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (t *Topic) GetTopicDetailsById(ctx context.Context, req *rpcModels.GetTopicParam, reply *rpcModels.GetTopicDetailsReply) error {
	topic, err := service.GetTopicById(req.Id)
	if err != nil {
		log.Error("service层<获取指定文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.TopicInfo.Id = topic.Id
	reply.TopicInfo.Title = topic.Title
	reply.TopicInfo.Author = topic.Author
	reply.TopicInfo.CategoryId = topic.CategoryId
	reply.TopicInfo.Updated = topic.Updated
	reply.TopicInfo.ReplayCount = topic.ReplayCount
	reply.TopicInfo.Uid = topic.Uid
	reply.TopicInfo.CategoryTitle = topic.Category.Title
	reply.TopicInfo.CategoryId = topic.Category.Id
	reply.Content = topic.Content

	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (t *Topic) GetTopicById(ctx context.Context, req *rpcModels.GetTopicParam, reply *rpcModels.GetTopic) error {
	topic, err := service.GetTopicById(req.Id)
	if err != nil {
		log.Error("service层<获取指定文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.Id = topic.Id
	reply.Title = topic.Title
	reply.Attachment = topic.Attachmemt
	reply.Content = topic.Content
	reply.Author = topic.Author
	reply.ResCode = ""
	reply.ResCode = ""
	return nil
}

func (t *Topic) GetTopicsByKeyword(ctx context.Context, req *rpcModels.SearTopicParams, reply *rpcModels.GetTopicsReply) error {
	infos, err := service.GetTopicsByKeyword(req.Keyword)
	if err != nil {
		log.Error("service层<搜索文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.List = *infos
	reply.Total = len(*infos)
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (t *Topic) AddTopic(ctx context.Context, req *rpcModels.TopicForm, reply *rpcModels.AddTopicReply) error {
	var topic = models.Topic{
		Uid: 0,
		Content: req.Content,
		Created: time.Now(),
		Title: req.Title,
		Author: req.Author,
		CategoryId: req.CatId,
		ReplayCount: 0,
		Attachmemt: req.Attachment,
	}
	var summary = models.TopicSummary{
		Title: req.Title,
		Content: req.Summary,
		Created: time.Now(),
		Author: req.Author,
	}
	err := service.AddTopic(topic, summary)
	if err != nil {
		log.Error("service层<添加文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.ResMsg = ""
	reply.ResCode = ""
	return nil
}

func (t *Topic) UpdateTopic(ctx context.Context, req *rpcModels.TopicForm, reply *rpcModels.UpdateTopicReply) error {
	err := service.UpdateTopic(*req)
	if err != nil {
		log.Error("service层<更新文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (t *Topic) GetTopicsByCatId(ctx context.Context, req *rpcModels.GetTopicsByCatParams, reply *rpcModels.GetSummarysReply) error {
	topics, err := service.GetTopicSummarysByCatId(req.CatId, req.PageNum, req.PageSize)
	if err != nil {
		log.Error("service层<搜索文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	total := service.GetTopicsCountByCatId(req.CatId)
	var list []rpcModels.TopicSummary
	for _,topic := range *topics {
		var info rpcModels.TopicSummary
		info.Id = topic.Id
		info.Author = topic.Author
		info.Title = topic.Title
		info.Created = topic.Created
		info.Content = topic.Content
		list = append(list, info)
	}
	reply.List = list
	reply.Total = total
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (t *Topic) GetTopics(ctx context.Context, req *rpcModels.GetTopicsParams, reply *rpcModels.GetTopicsReply) error {
	infos, err := service.GetTopics(req.PageNum, req.PageSize)
	if err != nil {
		log.Error("service层<获取指定文章>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	total := service.GetTopicsCount()
	reply.List = *infos
	reply.Total = total
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}