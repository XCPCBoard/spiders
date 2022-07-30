package scraper

import (
	"XCPCBoard/spiders/dao"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/5/12 13:36

type KV struct {
	Key string
	Val interface{}
}

//internalFlushRedis 内部刷新redis数据
func internalFlushRedis(kvs ...KV) {
	for _, kv := range kvs {
		// 底层库实现了自动重试
		err := dao.RedisClient.Set(kv.Key, kv.Val, 0).Err()
		if err != nil {
			log.Errorf("internal flush redis error %v", err)
		}
	}

}
