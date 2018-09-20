package utils

import (
	log "code.google.com/log4go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetConn() (*gorm.DB,error) {
	var args = GetConnUrl()
	db, err := gorm.Open(args[0], args[1])
	if err != nil {
		log.Error("数据库连接错误："+err.Error())
	}

	//db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
	return db,err
}

func GetConnUrl() []string {
	driverName := "mysql" //beego.AppConfig.String("driverName")
	username := "root" //beego.AppConfig.String("username")
	password := "root" //beego.AppConfig.String("password")
	host := "192.168.91.130" //beego.AppConfig.String("host")
	port := "3306"//beego.AppConfig.String("port")
	dbName := "blog" //beego.AppConfig.String("dbName")
	charset := "utf8" //beego.AppConfig.String("charset")
	url := ""
	// url := username+":"+password+"tcp@("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	if len(password) == 0 {
		url = username+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	} else {
		url = username+":"+password+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	}
	return []string{driverName, url}
}
