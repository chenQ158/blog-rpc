package utils

import (
	log "code.google.com/log4go"
	"github.com/go-redis/redis"
)

// 私有ClusterClient
var redisClusterClient *redis.ClusterClient
// 公有方法变量，使用闭包，用于获取ClusterClient
var REDIS_CLIENT = func() *redis.ClusterClient {return nil} ()

func GetRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:				"192.168.91.130:6379",
		Password:			"admin",
		DB:					0,
		PoolSize:			10,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error("Redis 连接失败：",err)
		return nil, err
	}
	return client, nil
}

func InitClusterClient() {
	var op = redis.ClusterOptions{
		Addrs:				[]string{"192.168.91.130:7001",
			"192.168.91.130:7002",
			"192.168.91.130:7003",
			"192.168.91.130:7004",
			"192.168.91.130:7005",
			"192.168.91.130:7006"},
		Password:			"",
		DialTimeout:		10,
		PoolSize:			10,
		ReadTimeout:		30,
		WriteTimeout:		30,
		PoolTimeout:		30,
	}

	//REDIS_CLIENT = redis.NewClusterClient(&op)
	REDIS_CLIENT = func() *redis.ClusterClient {
		// 获取op参数
		// 检查op参数并设置默认值
		return redis.NewClusterClient(&op)
	}()
	//REDIS_CLIENT.ConfigSet("cluster-enabled","yes")
}