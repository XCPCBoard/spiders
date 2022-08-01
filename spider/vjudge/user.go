package vjudge

import (
	"XCPCBoard/spiders/model"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/utils/keys"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

func vJudgeCallback(c *colly.Collector) {
	c.OnHTML("body", func(e *colly.HTMLElement) {
		uid := e.Request.Ctx.Get("uid")
		if uid == "" {
			log.Errorf("%v", model.UidError)
			return
		}
		// 最近24小时过题
		retStr := e.DOM.Find(".container a[title=\"New solved in last 24 hours\"]").First().Text()
		num, err := strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		} else {
			e.Request.Ctx.Put(keys.VJudgeLast1dayPassAmountKey(uid), num)
		}
		// 最近7天
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"New solved in last 7 days\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		} else {
			e.Request.Ctx.Put(keys.VJudgeLast7dayPassAmountKey(uid), num)
		}
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"New solved in last 30 days\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		} else {
			e.Request.Ctx.Put(keys.VJudgeLast30dayPassAmountKey(uid), num)
		}
		retStr = e.DOM.Find(fmt.Sprintf(".container a[title=\"Overall solved\"]")).First().Text()
		num, err = strconv.Atoi(retStr)
		if err != nil {
			log.Errorf("VJ strToInt get err:%v\t and the return is %v:", retStr, err)
		} else {
			e.Request.Ctx.Put(keys.VJudgePassAmountKey(uid), num)
		}
	})
}

func getPersonPage(uid string) string {
	return "https://vjudge.net/user/" + uid
}
