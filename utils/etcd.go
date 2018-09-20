package utils

import (
	"context"
	"github.com/coreos/etcd/client"
	transport2 "github.com/coreos/etcd/pkg/transport"
	log "code.google.com/log4go"
	"time"
)

var globalKApi client.KeysAPI
var globalEndpoints []string

func InitConfig(endpoints []string) {
	globalEndpoints = endpoints
}

func getKApi() client.KeysAPI {
	if globalKApi == nil {
		transport, err := transport2.NewTransport(transport2.TLSInfo{}, 30*time.Second)
		if err != nil {
			log.Debug(err)
		}
		cfg := client.Config{
			Endpoints: globalEndpoints,
			Transport: transport,
			HeaderTimeoutPerRequest: time.Second,
		}
		c, err := client.New(cfg)
		if err != nil {
			log.Debug(err)
		}
		globalKApi = client.NewKeysAPI(c)
	}
	return globalKApi
}

func GetKeyList(path string) (*client.Response, error) {
	kapi := getKApi()

	opts := &client.GetOptions{Recursive: true}
	return kapi.Get(context.Background(), path, opts)
}

func UpdateKey(key, value string) (*client.Response, error) {
	kapi := getKApi()
	return kapi.Update(context.Background(), key, value)
}

func SetKey(key, value string, opts *client.SetOptions) (*client.Response, error) {
	kapi := getKApi()
	return kapi.Set(context.Background(), key, value, opts)
}

func DeleteKey(key string, opts *client.DeleteOptions) (*client.Response, error) {
	kapi := getKApi()
	return kapi.Delete(context.Background(), key, opts)
}

func TestEtcd() {
	//cfg := client.Config{
	//	Endpoints:[]string{"http://192.168.92.128:2379"},
	//	Transport: client.DefaultTransport,
	//}
	//
	//c, err := client.New(cfg)
	//if err != nil {
	//	log.Error("Etcd client 连接失败：" + err.Error())
	//	return
	//}
	//
	//kAPI := client.NewKeysAPI(c)
	//// create a new key /foo with the value "bar"
	//_,err = kAPI.Create(context.Background(), "/foo", "bar")
	//if err != nil {
	//	log.Error("Etcd Create 失败：" + err.Error())
	//	return
	//}
	//
	//// delete the newly created key only if the value is still "bar"
	//_, err = kAPI.Delete(context.Background(), "/foo", &client.DeleteOptions{PrevValue:"bar"})
	//if err != nil {
	//	log.Error("Etcd Delete 失败：" + err.Error())
	//	return
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//// set a new key ,ignoring its previous state
	//_, err = kAPI.Set(ctx, "/ping", "pong", nil)
	//if err != nil {
	//	if err == context.DeadlineExceeded {
	//		log.Error("请求超时："+err.Error())
	//		return
	//	} else {
	//		log.Error("Etcd Set 失败："+err.Error())
	//		return
	//	}
	//}
}
