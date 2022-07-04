package main

import (
	_ "XCPCer_board/config"
	"XCPCer_board/dao"
	_ "XCPCer_board/dao"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	
}

func init() {
	redisClient, err := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	dbClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
	dao.DBClient = dbClient
}
