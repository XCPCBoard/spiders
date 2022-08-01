package nowcoder

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"

	"XCPCBoard/spiders/model"
	"XCPCBoard/utils/keys"
	log "github.com/sirupsen/logrus"
)

// @Author: Feng
// @Date: 2022/4/11 16:17

//-------------------------------------------------------------------------------------------//
// 基础方法
//-------------------------------------------------------------------------------------------//

const kScraperName = ""

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
				e.Request.Ctx.Put(keys.NowcoderRatingKey(uid), num)
			}
			// 排名
			num, err = strconv.Atoi(e.DOM.Find(getNowCoderContestBaseFindRule(ratingRankingKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(keys.NowcoderRankingKey(uid), num)
			}
			// 过题数
			num, err = strconv.Atoi(e.DOM.Find(getNowCoderContestBaseFindRule(contestAmountKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(keys.NowcoderContestAmountKey(uid), num)
			}
		},
	)

}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func (n *nowCoder) fetchMainPage(ctx *colly.Context) error {
	// 构造上下文，及传入参数
	uid := ctx.Get("uid")
	// 请求
	err := n.mainPage.Request("GET", getContestProfileUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return err
	}
	return nil
}
