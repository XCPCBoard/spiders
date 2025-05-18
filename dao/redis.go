package dao

import (
	"XCPCer_board/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var RedisClient *redis.Client

const redisDriver = "redis"

// NewRedisClient 初始化redis连接
func NewRedisClient() (*redis.Client, error) {
	// 获取配置
	redisConfig := config.Conf.Storages[redisDriver]

	// 初始化
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port), //"redis:6379"
		Password: redisConfig.Password,
		DB:       0, // use default DB
	})
	// 添加调试日志
	log.Infof("Redis配置: %+v", redisConfig)

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("debug:%v", redisConfig.Host)
		log.Errorf("Open Redis Error1:%v", err)
		return nil, err
	}
	return redisClient, nil
}
