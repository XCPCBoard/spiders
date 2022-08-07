package codeforces

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/spiders/scraper"
	"XCPCBoard/utils/keys"
)

func init() {
	scraper.GetStrategyInstance().Register(keys.CodeforcesKey, &codeforces{})
}

var (
	// 爬取函数
	fetchers = []func(codeforces, *colly.Context) error{
		codeforces.fetchUserInfo,
		codeforces.fetchAcceptInfo,
	}
)

type codeforces struct {
	problemCollector *colly.Collector
	infoCollector    *colly.Collector
}

func (c *codeforces) Init() {
	c.problemCollector = scraper.NewBaseCollector()
	problemCallback(c.problemCollector)
	c.infoCollector = scraper.NewBaseCollector()
	userInfoCallback(c.infoCollector)
}

//Scrape 拉取牛客的所有结果
func (c *codeforces) Scrape(ctx *colly.Context) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		err := f(*c, ctx)
		if err != nil {
			log.Errorf("do Fetcher Error %v", err)
		}
	}
}
