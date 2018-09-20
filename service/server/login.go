package server

import (
	"blog-rpc/rpc/models"
	"blog-rpc/service"
	log "code.google.com/log4go"
	"context"
	_ "github.com/go-sql-driver/mysql"
)

type Login struct {}

func (l *Login) Dologin(ctx context.Context, params *rpcModels.LoginParam, reply *rpcModels.LoginReply) error {
	log.Debug("service层<登录>方法request信息：%v", params)
	userptr, err := service.DoLogin(params.Username, params.Password)
	if err != nil {
		log.Error("service层<登录>方法出错：%v", err)
		//reply.ResCode =
		//reply.ResMsg =
		return err;
	}
	if userptr == nil {
		log.Debug("service层<登录>方法，用户名或密码不正确")
		reply.ResCode = "3001"
		reply.ResMsg = "用户名或密码错误"
		return nil
	}
	reply.Id = userptr.Id
	reply.Username = userptr.Username
	reply.Nickname = userptr.Nickname
	reply.ResCode = ""
	reply.ResMsg = ""
	log.Debug("service层<登录>方法请求：%v,响应：%v", params, reply)
	return nil
}