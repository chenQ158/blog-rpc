package service

import (
	"blog-rpc/error"
	"blog-rpc/models"
	"blog-rpc/utils"
	log "code.google.com/log4go"
	"github.com/jinzhu/gorm"
	"time"
)

func AddCat(name string) *errors.CommonError {
	db,err := utils.GetConn()
	if err != nil {
		log.Error("<添加分类>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	cat := models.Category{Title:name, Created:time.Now(), TopicTime:time.Now(), TopicCount: 0}
	err = db.Table("category").Save(&cat).Error
	if err != nil {
		log.Error("<添加分类>插入分类失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return nil
}

func GetCats(page,limit int) (*[]models.Category, *errors.CommonError) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		log.Error("<获取分类分页>数据库连接错误:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	err = db.Table("category").Limit(limit).Offset((page-1)*limit).Find(&cats).Error
	if err != nil {
		log.Error("<获取分类分页>查询分类失败:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &cats, nil
}

func GetAllCats() (*[]models.Category, *errors.CommonError) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		log.Error("<获取所有分类>数据库连接错误:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	err = db.Table("category").Find(&cats).Error
	if err != nil {
		log.Error("<获取所有分类>查询分类失败:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &cats, nil
}

func GetCatCount() (int, *errors.CommonError) {
	var count int
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<获取分类总数>数据库连接错误:%v", err)
		return count, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()

	err = db.Table("category").Count(&count).Error
	if err != nil {
		log.Error("<获取分类总数>查询分类失败:%v", err)
		return count, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return count, nil
}

func DelCat(id int) *errors.CommonError {
	db, err := utils.GetConn()
	if err != nil {
		log.Error("<删除分类>数据库连接错误:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	conn := db.Begin()
	err = conn.Table("category").Delete(models.Category{Id:id}).Error
	if err != nil {
		conn.Callback()
		log.Error("<删除分类>删除分类失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	conn.Table("topic").Where("CATEGORY_ID=?", id).Update(&models.Topic{CategoryId: 0})
	if err != nil {
		conn.Callback()
		log.Error("<删除分类>更新分类下文章失败:%v", err)
		return &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	conn.Commit()
	return nil
}

func GetCatsByKey(keyword string) (*[]models.Category, *errors.CommonError) {
	db, err := utils.GetConn()
	var cats []models.Category
	if err != nil {
		log.Error("<搜索分类>数据库连接错误:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	err = db.Table("category").Where("TITLE like ?", keyword+"%").Find(&cats).Error
	if err != nil {
		log.Error("<删除分类>更新分类下文章失败:%v", err)
		return &cats, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &cats, nil
}

func UpdateCatTopicCount(db *gorm.DB, catId int, isAdd bool) error {
	var err error
	if isAdd {
		err = db.Exec("UPDATE category SET TOPIC_COUNT = TOPIC_COUNT + 1 WHERE ID=?",catId).Error
	} else {
		err = db.Exec("UPDATE category SET TOPIC_COUNT = TOPIC_COUNT - 1 WHERE ID=? and TOPIC_COUNT > 0",catId).Error
	}
	if err != nil {
		log.Error("<更新分类下文章数>更新分类下文章数失败:%v", err)
		return err
	}
	return nil
}

func GetCatById(id int) (*models.Category, *errors.CommonError) {
	db, err := utils.GetConn()
	var cat models.Category
	if err != nil {
		log.Error("<搜索分类>数据库连接错误:%v", err)
		return &cat, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	defer db.Close()
	err = db.Table("category").Where("ID = ?", id).Find(&cat).Error
	if err != nil {
		log.Error("<删除分类>更新分类下文章失败:%v", err)
		return &cat, &errors.CommonError{ErrCode: "2002", ErrMsg: "数据中心异常！"}
	}
	return &cat, nil
}

