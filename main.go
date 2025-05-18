package main

import (
	_ "XCPCer_board/config"
	"XCPCer_board/dao"
	_ "XCPCer_board/dao"
	"XCPCer_board/spider/codeforces"
	"XCPCer_board/spider/luogu"
	"XCPCer_board/spider/nowcoder"
	"XCPCer_board/spider/vjudge"
	"database/sql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	//"XCPCer_board/spider/nowcoder"
	_ "github.com/FengZhg/go_tools/gin_logrus"
)

// 主入口函数
func main() {
	http.Handle("/metrics", promhttp.Handler())
	//启动 web 服务
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(2116), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Info("监控启动，端口为：" + strconv.Itoa(2116))
	}()

	c := cron.New()
	c.AddFunc("@every 120s", func() {
		log.Infoln("start scraper...")
		ls, err := dao.DBClient.Query("select uid,platform from id_platform;")
		defer func(ls *sql.Rows) {
			err := ls.Close()
			if err != nil {

			}
		}(ls)
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
			log.Infoln(platform)
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
