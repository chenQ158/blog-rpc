package service

import (
	"blog-rpc/error"
	"blog-rpc/models"
	"blog-rpc/rpc/models"
	"blog-rpc/utils"
	log "code.google.com/log4go"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

func DelTopic(id, categoryId int) *errors.CommonError {
	db,err := utils.GetConn()
	if err != nil {
		log.Error("<通过Id删除文章>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	conn := db.Begin()
	//db.Where("id=?", id).Delete(models.Topic{})
	var topic = models.Topic{Id:id}
	err = conn.Table("topic").Delete(&topic).Error
	if err != nil {
		conn.Rollback()
		log.Error("<通过Id删除文章>删除文章失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	err = conn.Table("topic_summary").Where("ID=?", topic.Id).Delete(&models.TopicSummary{}).Error
	if err != nil {
		conn.Rollback()
		log.Error("<通过Id删除文章>删除文章概述失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	err = UpdateCatTopicCount(conn, categoryId, false)
	if err != nil {
		conn.Rollback()
		log.Error("<通过Id删除文章>删除文章失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	conn.Commit()
	return nil
}

func GetTopics(page, limit int) (*[]rpcModels.TopicInfo, *errors.CommonError) {
	db, err := utils.GetConn()
	var infoList []rpcModels.TopicInfo
	if err != nil {
		log.Error("<获取文章分页>数据库连接错误:%v", err)
		return &infoList, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	//fmt.Println(page,)
	nativeSql := "select topic.ID ID,topic.TITLE TITLE,topic.ATTACHMENT ATTACHMENT,topic.REPLAY_COUNT REPLAY_COUNT,topic.UPDATED UPDATED,user.ID UID,user.USERNAME AUTHOR,category.ID CATEGORY_ID,category.TITLE CATEGORY_TITLE from topic left join category on topic.CATEGORY_ID=category.ID left join user on user.ID=topic.USER_ID limit ?, ?"
	resList := GeneralQuery(db, nativeSql, (page-1)*limit, limit)
	//fmt.Println("resList:", resList)
	btArr, err := json.Marshal(resList)
	// fmt.Println("GetTopicsByKeyword btArr:", string(btArr))
	if err != nil {
		log.Error("<获取文章分页>文章详情查询结果转json失败:%v", err)
		return &infoList, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}

	err = json.Unmarshal(btArr, &infoList)
	if err != nil {
		log.Error("<获取文章分页>json转文章详情失败:%v", err)
		return nil, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &infoList, nil
}

func GetTopicById(id int) (*models.Topic, *errors.CommonError) {
	db, err := utils.GetConn()
	var topic models.Topic
	if err != nil {
		log.Error("<获取指定文章>数据库连接错误:%v", err)
		return &topic, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	var cat models.Category
	err = db.Table("topic").First(&topic, "id=?", id).Error
	if err != nil {
		log.Error("<获取指定文章>查询文章失败:%v", err)
		return nil, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	db.Table("category").Where("ID=?", topic.CategoryId).First(&cat)
	topic.Category = cat
	return &topic, nil
}

func AddTopic(topic models.Topic, summary models.TopicSummary) *errors.CommonError {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<添加文章>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	topic.Updated = time.Now()
	topic.Created = time.Now()
	topic.ReplayTime = time.Now()
	conn := db.Begin()
	err = conn.Table("topic").Save(&topic).Error
	if err != nil {
		conn.Rollback()
		log.Error("<添加文章>插入文章失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}

	summary.Id = topic.Id
	err = conn.Table("topic_summary").Save(&summary).Error
	if err != nil {
		conn.Rollback()
		log.Error("<添加文章>添加文章概述失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	err = UpdateCatTopicCount(conn, topic.CategoryId, true)
	if err != nil {
		conn.Rollback()
		log.Error("<添加文章>更新分类文章数失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	conn.Commit()
	return nil
}

func UpdateTopic(form rpcModels.TopicForm) *errors.CommonError {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<更新文章>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	var topic = models.Topic{Title:form.Title,Content:form.Content,Attachmemt:form.Attachment}
	err = db.Table("topic").Where("Id=?", form.Id).Updates(topic).Error
	if err != nil {
		log.Error("<更新文章>更新文章失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return nil
}

func GetSummarysByOffsetAndLimit(page, limit int) (*[]models.TopicSummary, *errors.CommonError) {
	var topics []models.TopicSummary
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取文章分页>数据库连接错误:%v", err)
		return &topics, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	err = db.Table("topic_summary").Limit(limit).Offset((page-1)*limit).Find(&topics).Error
	//err = db.Table("topic").Limit(limit).Offset((page-1)*limit).Find(&topics).Error
	if err != nil {
		log.Error("<获取文章分页>获取文章分页失败:%v", err)
		return &topics, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &topics,nil
}

func GetTopicsCount() int {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取文章总数>数据库连接错误:%v", err)
		return 0
	}
	defer db.Close()
	var count int
	db.Table("topic").Count(&count)
	return count
}

func GetTopicSummarysByCatId(catId, pageNum, limit int) (*[]models.TopicSummary, *errors.CommonError) {
	var topics []models.TopicSummary
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取分类下文章>数据库连接错误:%v", err)
		return &topics, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	err = db.Table("topic_summary").
		Select("topic_summary.ID, topic_summary.AUTHOR,topic_summary.TITLE,topic_summary.CONTENT,topic_summary.CREATED").
		Joins("join topic on topic_summary.ID=topic.ID").
		Where("topic.CATEGORY_ID=?", catId).Limit(limit).Offset((pageNum-1)*limit).Find(&topics).Error
	if err != nil {
		log.Error("<获取分类下文章>查询分类下文章失败:%v", err)
		return &topics, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &topics,nil
}

func GetTopicsCountByCatId(catId int) int {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取分类下文章数量>数据库连接错误:%v", err)
		return 0
	}
	defer db.Close()

	var count int
	db.Table("topic").Where("CATEGORY_ID=?", catId).Count(&count)
	return count
}

func UpdateTopicCommentCount(db *gorm.DB, id int, isAdd bool) error {
	var err error
	if isAdd {
		err = db.Exec("UPDATE topic SET REPLAY_COUNT = REPLAY_COUNT + 1 WHERE ID=?",id).Error
	} else {
		err = db.Exec("UPDATE topic SET REPLAY_COUNT = REPLAY_COUNT - 1 WHERE ID=? and REPLAY_COUNT > 0",id).Error
	}
	if err != nil {
		log.Error("<更新文章评论数>更新文章评论数失败:%v", err)
		return err
	}
	return nil
}


func GetTopicsByKeyword(keyword string) (*[]rpcModels.TopicInfo, *errors.CommonError) {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<搜索文章标题>数据库连接错误："+err.Error())
		return nil, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	nativeSql := "select topic.ID ID,topic.TITLE TITLE,topic.ATTACHMENT ATTACHMENT,topic.REPLAY_COUNT REPLAY_COUNT,topic.UPDATED UPDATED,user.ID UID,user.USERNAME AUTHOR,category.ID CATEGORY_ID,category.TITLE CATEGORY_TITLE from topic left join category on topic.CATEGORY_ID=category.ID left join user on user.ID=topic.USER_ID where topic.TITLE like ?"
	resList := GeneralQuery(db, nativeSql, keyword+"%")
	//fmt.Println("resList:", resList)
	btArr, err := json.Marshal(resList)
	//fmt.Println("GetTopicsByKeyword btArr:", string(btArr))
	if err != nil {
		log.Error("<搜索文章标题>文章详情查询结果转json失败")
		return nil, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	var infoList []rpcModels.TopicInfo
	err = json.Unmarshal(btArr, &infoList)
	if err != nil {
		log.Error("<搜索文章标题>json转文章详情失败")
		return nil, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	//fmt.Println("InfoList:", infoList)
	return &infoList, nil
}

// select topic.ID,topic.TITLE,topic.ATTACHMENT,topic.UPDATED,topic.VIEWS,topic.AUTHOR,category.TITLE,user.USERNAME from topic join category on category.ID = topic.CATEGORY_ID join user on user.ID = topic.USER_ID;
func GetTopicInfo(id int) (string, *errors.CommonError) {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取文章详情>数据库连接错误：%v", err)
		return "", &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	nativeSql := "select topic.ID,topic.TITLE,topic.ATTACHMENT,topic.UPDATED,topic.VIEWS,topic.AUTHOR,category.TITLE,user.USERNAME from topic join category on category.ID = topic.CATEGORY_ID join user on user.ID = topic.USER_ID where topic.ID = ?"
	rows, err := db.Raw(nativeSql, id).Rows()
	if err != nil {
		log.Error("<获取文章详情>查询文章详情失败")
		return "", &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}

	defer rows.Close()

	var ret []interface{}
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		rows.Scan(scanArgs...)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}
			}
		}
		ret = append(ret, record)
	}

	str, _ := json.Marshal(ret)
	return string(str), nil
}