package vjudge

import (
	"context"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/spiders/dao"
)

//SetUserMsgToRedis 将用户信息放入redis
func SetUserMsgToRedis(uid string, ctx context.Context) error {
	//get user msg
	res, err := ScrapeUser(uid)
	if err != nil {
		log.Errorf("%v get uid=%v message err:%v", packageName, uid, err)
	}
	//set data to redis
	for key, val := range mapKey {
		err := dao.RedisClient.Set(ctx, key, val, 0).Err()
		if err != nil {
			log.Errorf("%v set redis data for uid=%v failed, err:%v\n", packageName, uid, err)
			return err
		}
	}
	return nil
}
