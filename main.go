package main

import (
	_ "XCPCer_board/config"
	"XCPCer_board/dao"
	_ "XCPCer_board/dao"
	"XCPCer_board/spider/codeforces"
	"XCPCer_board/spider/luogu"
	"XCPCer_board/spider/nowcoder"
	"XCPCer_board/spider/vjudge"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"

	//"XCPCer_board/spider/nowcoder"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	//luogu.Luogu("chenyyy28")
	c := cron.New()
	c.AddFunc("@every 180s", func() {
		log.Infoln("okk")
		ls, err := dao.DBClient.Query("select uid,platform from id_platform;")
		defer ls.Close()
		if err != nil {
			log.Errorf("database error: %v", err)
			return
		}
		for ls.Next() {
			id, platform := "", ""
			err := ls.Scan(&id, &platform)
			if err != nil {
				log.Errorln(err)
				return
			}
			if platform == "codeforces" {
				codeforces.Flush(id)
			} else if platform == "nowcoder" {
				nowcoder.Flush(id)
			} else if platform == "vjudge" {
				vjudge.Flush(id)
			} else if platform == "luogu" {
				luogu.Luogu(id)
			}

		}
	})
	c.Start()
	go c.Start()
	defer c.Stop()

	select {}

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
