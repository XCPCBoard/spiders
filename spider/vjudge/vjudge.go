package vjudge

import (
	"XCPCBoard/spiders/scraper"
	"github.com/gocolly/colly"
)

type vJudge struct {
	collector *colly.Collector
}

func (v *vJudge) Init() {
	v.collector = scraper.NewBaseCollector()
	vJudgeCallback(v.collector)
}

func (v *vJudge) Scrape(ctx *colly.Context) {
	// 构造上下文，及传入参数
	uid := ctx.Get("uid")
	// 请求
	err := v.collector.Request("GET", getPersonPage(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("scraper error %v", err)
	}
}
