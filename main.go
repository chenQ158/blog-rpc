package main

import (
	server2 "blog-rpc/service/server"
	"blog-rpc/utils"
	log "code.google.com/log4go"
	"github.com/go-metrics-master"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
	"os"
	"os/signal"
	"strings"
	"time"
)

var endpoints = "http://192.168.91.130:2379"
//var addr = "127.0.0.1:8974"
var addr = "127.0.0.1:8975"
var basepath = "/rpcx"

func main() {

	// 初始化etcd配置
	utils.InitConfig(strings.Split(endpoints, ","))

	server := rpcx.NewServer()
	rplugin := &plugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		EtcdServers: strings.Split(endpoints, ","),
		BasePath: basepath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: 5*time.Minute,
	}

	rplugin.Start()
	log.Debug("The EtcdRegisterPlugin start success")
	server.PluginContainer.Add(rplugin)
	server.PluginContainer.Add(plugin.NewMetricsPlugin())
	// 注册rpc服务
	loginServiceName := "rpc.Login"
	server.RegisterName(loginServiceName, new(server2.Login), "weight=1&m=devops")
	catgoryServiceName := "rpc.Category"
	server.RegisterName(catgoryServiceName, new(server2.Category), "weight=1&m=devops")
	topicServiceName := "rpc.Topic"
	server.RegisterName(topicServiceName, new(server2.Topic), "weight=1&m=devops")
	commentServiceName := "rpc.Comment"
	server.RegisterName(commentServiceName, new(server2.Comment), "weight=1&m=devops")

	// 如果系统异常退出，则注销服务
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		si := <-c
		if si != nil {
			rplugin.UnregisterAll()
			os.Exit(0)
		}
	}()

	// 监听服务
	server.Serve("tcp", addr)
}