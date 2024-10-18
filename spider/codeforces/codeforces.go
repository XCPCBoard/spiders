package codeforces

import (
	"XCPCer_board/dao"
	"XCPCer_board/scraper"
	log "github.com/sirupsen/logrus"
)

var (
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV, error){
		fetchUserInfo,
		fetchAcceptInfo,
	}
)

// scrape 拉取cf的所有结果
func scrape(uid string) (res []scraper.KV) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		kvs, err := f(uid)
		if err != nil {
			log.Errorf("do Fetcher Error %v", err)
			continue
		}
		res = append(res, kvs...)
	}
	return res
}

// Flush 刷新某用户cf-id信息
func Flush(uid string) {
	uuid = uid
	// 拉出所有kv对
	kvs := scrape(uid)

	name_id := 0
	if kvs == nil {
		log.Errorf("kv nil")
		return
	}
	//检测是否存在对应的成员
	res, err := dao.DBClient.Query("select (name_id)from id_platform where uid = ?&&platform=?", uid, "codeforces")
	if err != nil {
		log.Errorf("sql error %v", err)
	} else {
		for res.Next() {
			res.Scan(&name_id)
		}
		if name_id == 0 {
			log.Errorf("null name_id ,cant find name_id")
		}
	}
	//更新数据库信息
	for _, j := range kvs {
		if j.Key == "gxuicpc:codeforces_problem_pass_"+uid {
			scraper.FlushDB("update  score set cf_problem = ? where id = ?;", j.Val, name_id)
		}
		if j.Key == "gxuicpc:codeforces_max_rating_"+uid {
			scraper.FlushDB("update score set cf_rank = ? where id = ?;", j.Val, name_id)
		}
	}
	//更新redis信息
	scraper.FlushRedis(kvs)
	// 向持久化处理协程注册持久化处理函数
	scraper.CustomFlush(func() error {
		log.Infoln(kvs)
		return nil
	})
}
