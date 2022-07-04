package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var (
	// 爬取函数
	fetchers = []func(uid string) ([]scraper.KV, error){
		fetchMainPage,
		fetchConPage,
	}
	// 匹配持久化处理函数
	persistHandlerMap = map[string]func(uid string) func(string, interface{}) error{
		RatingKey:     profilePersistHandler,
		contestSumKey: profilePersistHandler,
		rankKey:       profilePersistHandler,
		submissionKey: submissionPersistHandler,
	}
)

//scrape 拉取atCoder的所有结果
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
	fmt.Println(len(res))
	return res
}

//profilePersistHandler 个人主页持久化函数
func profilePersistHandler(uid string) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		//dao.RedisClient.Set()
		//dao.DBClient.ExecContext()
		log.Infof("atcoder uid :%v Key %v Val %v", uid, key, val)
		return nil
	}
}

//submissionPersistHandler submission持久化函数
func submissionPersistHandler(uid string) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		//dao.RedisClient.Set()
		//dao.DBClient.ExecContext()
		log.Infof("atcoder uid :%v Key %v Val %v", uid, key, val)
		return nil
	}
}

//matchPersistHandlers 匹配持久化函数
func matchPersistHandlers(uid string, kvs []scraper.KV) []scraper.Persist {
	var res []scraper.Persist
	for ind, _ := range kvs {
		h, ok := persistHandlerMap[kvs[ind].Key]
		if ok {
			res = append(res, kvs[ind].GetPersistHandler(scraper.NewPersistHandler(h(uid))))
		}
	}
	return res
}

//Flush 刷新atCoder某用户信息
func Flush(uid string) {
	// 拉出所有kv对
	kvs := scrape(uid)
	// 为所有key对匹配持久化函数
	persists := matchPersistHandlers(uid, kvs)
	// 向持久化处理协程注册持久化处理函数
	scraper.RegisterPersist(persists...)
}
