package service

import (
	"blog-rpc/error"
	"blog-rpc/models"
	"blog-rpc/utils"
	log "code.google.com/log4go"
)

func AddComment(comment *models.Comment) *errors.CommonError {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<添加评论>数据库连接异常：%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	conn := db.Begin()
	log.Debug("<添加评论>插入的评论信息：%v", comment)
	err = conn.Table("comment").Create(comment).Error
	if err != nil {
		conn.Rollback()
		log.Error("<添加评论>插入评论失败：%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	err = UpdateTopicCommentCount(conn, comment.TopicId, true)
	if err != nil {
		conn.Rollback()
		log.Error("<添加评论>更新评论数失败：%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	conn.Commit()
	log.Debug("<添加评论>插入的评论为:%v", comment)
	return nil
}

func GetCommentsByTopicId(id int) (*[]models.Comment, *errors.CommonError) {
	var comments []models.Comment
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取文章评论>数据库连接错误：%v", err)
		return &comments, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	log.Debug("<获取文章评论>查询的条件:%v",id)
	err = db.Table("comment").Where("TopicId=?",id).Find(&comments).Error
	if err != nil {
		log.Error("<获取文章评论>评论查询失败:%v", err)
		return &comments, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &comments, nil
}

func DelCommentById(id, topicId int) *errors.CommonError {
	var comment models.Comment
	var err error
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<通过Id删除评论>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	conn := db.Begin()

	err = conn.Table("comment").Where("ID=?", id).Delete(&comment).Error
	if err != nil{
		log.Error("<通过Id删除评论>删除评论失败:%v", err)
		conn.Rollback()
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	err = UpdateTopicCommentCount(conn, comment.TopicId, false)
	if err != nil {
		log.Error("<通过Id删除评论>更新文章评论数失败:%v", err)
		conn.Rollback()
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}

	conn.Commit()
	return nil
}