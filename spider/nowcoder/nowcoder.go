package nowcoder

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"

	"XCPCBoard/spiders/scraper"
	"XCPCBoard/utils/keys"
)

// @Author: Feng
// @Date: 2022/4/8 17:09

func init() {
	scraper.GetStrategyInstance().Register(keys.NowcoderKey, &NowCoder{})
}

var (
	// 爬取函数
	fetchers = []func(NowCoder, *colly.Context) error{
		NowCoder.fetchMainPage,
		NowCoder.fetchPractice,
	}
)

type NowCoder struct {
	mainPage     *colly.Collector
	practicePage *colly.Collector
}

func (n *NowCoder) Init() {
	n.mainPage = scraper.NewBaseCollector()
	enrichMainPageCollector(n.mainPage)
	n.practicePage = scraper.NewBaseCollector()
	enrichPracticePageCollector(n.practicePage)
}

//Scrape 拉取牛客的所有结果
func (n *NowCoder) Scrape(ctx *colly.Context) {
	// 请求所有
	for _, f := range fetchers {
		// 请求
		err := f(*n, ctx)
		if err != nil {
			log.Errorf("GetPersistHandler Fetcher Error %v", err)
		}
	}
}
