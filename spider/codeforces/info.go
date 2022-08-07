package codeforces

import (
	"encoding/json"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/spiders/model"
	"XCPCBoard/utils/keys"
)

// @Author: Feng
// @Date: 2022/5/12 22:41

//userInfoCallback 处理codeforces的api
func userInfoCallback(c *colly.Collector) {
	c.OnScraped(func(r *colly.Response) {
		// 获取uid
		uid := r.Request.Ctx.Get("uid")
		if uid == "" {
			log.Errorf("%v", model.UidError)
			return
		}
		// 反序列化
		rsp := &UserInfo{}
		err := json.Unmarshal(r.Body, rsp)
		if err != nil {
			log.Errorf("Codeforces User Info Unmarshal Error %v", err)
			return
		}
		if rsp.GetStatus() != "OK" || len(rsp.GetResult()) != 1 {
			log.Errorf("Response: %v Infos Length: %v", rsp.GetStatus(), len(rsp.GetResult()))
			return
		}
		info := rsp.GetResult()[0]
		if info.GetRating() != 0 {
			r.Ctx.Put(keys.CodeforcesRatingKey(uid), info.GetRating())
		}
		if info.GetMaxRating() != 0 {
			r.Ctx.Put(keys.CodeforcesMaxRatingKey(uid), info.GetMaxRating())
		}
		if info.GetRank() != "" {
			r.Ctx.Put(keys.CodeforcesRankingKey(uid), info.GetRank())
		}
		if info.GetMaxRank() != "" {
			r.Ctx.Put(keys.CodeforcesMaxRankingKey(uid), info.GetMaxRank())
		}
	})
}

//---------------------------------------------------------------------//
// 对外暴露函数:用户信息获取
//---------------------------------------------------------------------//

//fetchUserInfo 抓取用户信息
func (c *codeforces) fetchUserInfo(ctx *colly.Context) error {
	// 构造上下文，及传入参数
	uid := ctx.Get("uid")
	// 请求
	err := c.infoCollector.Request("GET", getUserInfoUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return err
	}
	return nil
}
