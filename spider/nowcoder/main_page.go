package nowcoder

import (
	"XCPCBoard/spiders/model"
	"XCPCBoard/spiders/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// @Author: Feng
// @Date: 2022/4/11 16:17

//-------------------------------------------------------------------------------------------//
// 基础方法
//-------------------------------------------------------------------------------------------//

const kScraperName = ""

var (
	mainScraper = scraper.InitializeAndRegisterScraper(
		kScraperName,
		enrichMainPageCollector,
	)
)

//enrichMainPageCollector 处理牛客个人主页的回调函数
func enrichMainPageCollector(c *colly.Collector) {
	//用goquery
	c.OnHTML(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix .my-state-main",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			if uid == "" {
				log.Errorf("%v", model.UidError)
				return
			}
			// rating
			num, err := strconv.Atoi(e.DOM.Find(fmt.Sprintf(".my-state-item:contains(%v) .state-num.rate-score5",
				ratingKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(GetRatingKey(uid), num)
			}
			// 排名
			num, err = strconv.Atoi(e.DOM.Find(getNowCoderContestBaseFindRule(ratingRankingKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(GetRankingKey(uid), num)
			}
			// 过题数
			num, err = strconv.Atoi(e.DOM.Find(getNowCoderContestBaseFindRule(contestAmountKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(utils.GetContestAmountKey(uid), num)
			}
		},
	)

}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchMainPage(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	// 请求
	err := mainScraper.Collector.Request("GET", getContestProfileUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return nil, err
	}
	// 解构出kv对
	kvs := scraper.Parse(ctx, map[string]struct{}{
		"uid": {},
	})
	return kvs, nil
}
