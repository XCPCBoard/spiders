package dao

import (
	"XCPCBoard/spiders/config"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

const redisDriver = "redis"

//NewRedisClient 初始化redis连接
func NewRedisClient() (*redis.Client, error) {
	// 获取配置
	redisConfig := config.Conf.Storages[redisDriver]
	// 初始化
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Errorf("Open Redis Error:%v", err)
		return nil, err
	}
	return redisClient, nil
}
