package nowcoder

import (
	"XCPCer_board/dao"
	"XCPCer_board/scraper"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/8 17:09

var (
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV, error){
		fetchMainPage,
		fetchPractice,
	}
)

// scrape 拉取牛客的所有结果
func scrape(uid string) (res []scraper.KV) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		kvs, err := f(uid)
		if err != nil {
			log.Errorf("GetPersistHandler Fetcher Error %v", err)
			continue
		}
		res = append(res, kvs...)
	}
	return res
}

// Flush 刷新某用户牛客id信息
func Flush(uid string) {
	// 拉出所有kv对
	kvs := scrape(uid)
	name_id := 0
	if kvs == nil {
		log.Errorf("kv nil")
		return
	}
	res, err := dao.DBClient.Query("select (name_id)from id_platform where uid = ?&& platform=?", uid, "nowcoder")
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
	for _, j := range kvs {
		if j.Key == "gxuicpc:nowcoder_pass_amount_"+uid {
			scraper.FlushDB("update  score set nk_problem = ? where id = ?;", j.Val, name_id)
		}
		if j.Key == "gxuicpc:nowcoder_ranking_"+uid {
			scraper.FlushDB("update score set nk_rank = ? where id = ?;", j.Val, name_id)
		}
	}

	scraper.FlushRedis(kvs)
	scraper.CustomFlush(func() error {
		log.Infoln(kvs)
		//scraper.FlushRedis(kvs)
		return nil
	})
}
