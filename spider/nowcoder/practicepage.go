package nowcoder

import (
	"strconv"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/spiders/model"
	"XCPCBoard/utils/keys"
)

// @Author: Feng
// @Date: 2022/4/11 16:17

//-------------------------------------------------------------------------------------------//
// 基础方法
//-------------------------------------------------------------------------------------------//

//enrichPracticePageCollector 处理牛客个人练习页面的回调函数
func enrichPracticePageCollector(c *colly.Collector) {
	//用goquery
	c.OnHTML(".nk-container.acm-container .nk-container .nk-main.with-profile-menu.clearfix .my-state-main",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			if uid == "" {
				log.Errorf("%v", model.UidError)
				return
			}
			// 题目通过数量
			num, err := strconv.Atoi(e.DOM.Find(getNowCoderContestBaseFindRule(passAmountKeyWord)).First().Text())
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(keys.NowcoderPassAmountKey(uid), num)
			}
		},
	)
}

//---------------------------------------------------------------------//
// 对外暴露函数:个人练习信息获取
//---------------------------------------------------------------------//

//fetchPractice 抓取个人练习页面的所有
func (n *nowCoder) fetchPractice(ctx *colly.Context) error {
	// 从上下文拉取uid
	uid := ctx.Get("uid")
	// 请求
	err := n.practicePage.Request("GET", getContestPracticeUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return err
	}
	return nil
}
