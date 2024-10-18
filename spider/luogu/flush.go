package luogu

import (
	"XCPCer_board/dao"
	"XCPCer_board/scraper"
	"context"
	log "github.com/sirupsen/logrus"
)

func flushDB(uid string, kvs []scraper.KV) {
	var name_id int
	res, err := dao.DBClient.Query("select (name_id)from id_platform where uid = ?&&platform=?", uid, "luogu")
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
		if j.Key == getAmountPassKey(uid) {
			_, err := dao.DBClient.Exec("update score set luogu_problem = ? where id = ?;", j.Val, name_id)
			if err != nil {
				log.Errorf("update luogu problem error,cause: sql error %v", err)
			}
		}
	}
}

func flushRedis(uid string, kvs []scraper.KV) {
	for _, kv := range kvs {
		err := dao.RedisClient.Set(context.Background(), kv.Key, kv.Val, 0).Err()
		if err != nil {
			log.Errorf("internal flush redis error %v", err)
		}
	}
}

func Flush(uid string, kvs []scraper.KV) {
	flushDB(uid, kvs)
	flushRedis(uid, kvs)
}
