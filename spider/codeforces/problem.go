package codeforces

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var (
	//接收 函数NewScraper 返回值 &Scraper
	//函数problemCallback作为传递参数
	problemScraper = scraper.NewScraper(
		problemCallback,
	)
	uuid = ""
)

func problemCallback(c *colly.Collector) {
	//为html选择器注册回调函数
	c.OnHTML("div[style=\"position: relative;\"] #pageContent ._UserActivityFrame_frame "+
		".roundbox.userActivityRoundBox ._UserActivityFrame_footer ._UserActivityFrame_countersRow",
		func(e *colly.HTMLElement) {
			tmp := strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", lastMonthPassKeyWord)).First().Text(), " ")[0]
			// 最近一个月过题数
			num, err := strconv.Atoi(tmp)
			if err != nil {
				if tmp != "" {
					log.Errorf("str atoi Error %v", err)
				}
			} else {
				e.Request.Ctx.Put(lastMonthPassAmount+"_"+uuid, num)
			}

			// 总过题数
			tmp = strings.Split(e.DOM.Find(fmt.Sprintf("._UserActivityFrame_counter:contains(solved):contains("+
				"%v) ._UserActivityFrame_counterValue", problemPassKeyWord)).First().Text(), " ")[0]
			num, err = strconv.Atoi(tmp)
			if err != nil {
				if tmp != "" {
					log.Errorf("str atoi Error %v", err)
				}
			} else {
				e.Request.Ctx.Put(problemPassAmountKey+"_"+uuid, num)
			}
		},
	)
}

// fetchAcceptInfo 获取过题情况
func fetchAcceptInfo(uid string) ([]scraper.KV, error) {
	// 构造上下文，及传入参数
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	//hdr := http.Header{"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0"}}
	//ctx.Put("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0")
	// 请求
	err := problemScraper.C.Request("GET", getPersonPageUrl(uid), nil, ctx, nil)
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
