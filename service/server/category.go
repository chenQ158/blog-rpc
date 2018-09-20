package server

import (
	"blog-rpc/rpc/models"
	"blog-rpc/service"
	log "code.google.com/log4go"
	"context"
)

type Category struct {}

func (c *Category) AddCat(ctx context.Context, req *rpcModels.AddCatParams, reply *rpcModels.AddCatReply) error {
	err := service.AddCat(req.CatName)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<添加分类>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (c *Category) GetCats(ctx context.Context, req *rpcModels.GetCatsParams, reply *rpcModels.GetCatsReply) error {
	cats, err := service.GetCats(req.PageNum, req.PageSize)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<获取分类分页>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	total, err := service.GetCatCount()
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<获取分类分页>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}

	var list []rpcModels.CatInfo
	for _,cat := range *cats {
		var info rpcModels.CatInfo
		info.CatName = cat.Title
		info.Id = cat.Id
		info.Updated = cat.TopicTime
		info.TopicCount = cat.TopicCount
		list = append(list, info)
	}
	reply.List = list
	reply.Total = total
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (c *Category) GetAllCats(ctx context.Context, req *int, reply *rpcModels.GetAllCats) error {
	cats, err := service.GetAllCats()
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<获取所有分类>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	var list []rpcModels.SimpleCatInfo
	for _,cat := range *cats {
		var info rpcModels.SimpleCatInfo
		info.CatName = cat.Title
		info.Id = cat.Id
		list = append(list, info)
	}
	reply.List = list
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

func (c *Category) DelCat(ctx context.Context, req *rpcModels.DelCatParams, reply *rpcModels.DelCatReply) error {
	err := service.DelCat(req.Id)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<删除分类>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	reply.ResMsg = ""
	reply.ResCode = ""
	return nil
}

func (c *Category) GetCatsByKey(ctx context.Context, req *rpcModels.SearchCatParams, reply *rpcModels.GetCatsReply) error {
	cats, err := service.GetCatsByKey(req.Keyword)
	if err != nil {
		// 根据返回的错误编码判断
		log.Error("service层<搜索分类>方法调用失败:%v", err)
		reply.ResCode = err.ErrCode
		reply.ResMsg = err.ErrMsg
		return nil
	}
	var list []rpcModels.CatInfo
	for _,cat := range *cats {
		var info rpcModels.CatInfo
		info.CatName = cat.Title
		info.Id = cat.Id
		info.Updated = cat.TopicTime
		info.TopicCount = cat.TopicCount
		list = append(list, info)
	}
	reply.List = list
	reply.Total = len(list)
	reply.ResCode = ""
	reply.ResMsg = ""
	return nil
}

