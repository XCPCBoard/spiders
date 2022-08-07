package codeforces

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/utils/keys"
)

func problemCallback(c *colly.Collector) {
	c.OnHTML("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame "+
		".roundbox.userActivityRoundBox ._UserActivityFrame_footer ._UserActivityFrame_countersRow",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			// 最近一个月过题数
			num, err := strconv.Atoi(strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", lastMonthPassKeyWord)).First().Text(), " ")[0])
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(keys.CodeforcesLastMonthPassAmountKey(uid), num)
			}
			// 总过题数
			num, err = strconv.Atoi(strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", problemPassKeyWord)).First().Text(), " ")[0])
			if err != nil {
				log.Errorf("str atoi Error %v", err)
			} else {
				e.Request.Ctx.Put(keys.CodeforcesPassAmountKey(uid), num)
			}
		},
	)
}

//fetchAcceptInfo 获取过题情况
func (c *codeforces) fetchAcceptInfo(ctx *colly.Context) error {
	// 构造上下文，及传入参数
	uid := ctx.Get("uid")
	// 请求
	err := c.problemCollector.Request("GET", getPersonPageUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
		return err
	}
	return nil
}
