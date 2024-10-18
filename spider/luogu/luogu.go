package luogu

import (
	"XCPCer_board/scraper"
	log "github.com/sirupsen/logrus"
)

var (
	fetchers = []func(uid string) []scraper.KV{
		GetProblemKvs,
	}
)

func scrape(uid string) (res []scraper.KV) {
	for _, f := range fetchers {
		kvs := f(uid)
		res = append(res, kvs...)
	}
	return res
}

func Luogu(uid string) {
	/**
	Todo
	希望在提交列表，通过用户名，获取userid，但是该页面需登录，懒了不想写登录
	*/
	kvs := scrape(uid)

	if kvs == nil {
		log.Errorf("kvs is empty")
	}

	Flush(uid, kvs)
	log.Infoln(kvs)
}
