package utils

import (
	log "code.google.com/log4go"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConn() (*gorm.DB,error) {
	var args []string = GetConnUrl()
	db, err := gorm.Open(args[0], args[1])
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
	}

	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	return db,err
}

func GetConnUrl() []string {
	driverName := beego.AppConfig.String("driverName")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbName := beego.AppConfig.String("dbName")
	charset := beego.AppConfig.String("charset")
	url := ""
	// url := username+":"+password+"tcp@("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	if len(password) == 0 {
		url = username+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	} else {
		url = username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	}
	return []string{driverName, url}
}
