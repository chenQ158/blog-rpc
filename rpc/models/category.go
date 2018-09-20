package rpcModels

import "time"

// 添加分类
type AddCatParams struct {
	CatName			string
}

type AddCatReply struct {
	baseReply
}

// 获取分类分页
type GetCatsParams struct {
	PageNum			int
	PageSize		int
}

type CatInfo struct {
	Id				int
	CatName			string
	TopicCount		int
	Updated			time.Time
}

type GetCatsReply struct {
	baseReply
	List			[]CatInfo
	Total			int
}

type SimpleCatInfo struct {
	Id				int
	CatName			string
}

type GetAllCats struct {
	baseReply
	List			[]SimpleCatInfo
}

type SearchCatParams struct {
	Keyword			string
}

// 删除分类
type DelCatParams struct {
	Id  			int
}

type DelCatReply struct {
	baseReply
}