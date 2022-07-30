package scraper

import (
	"github.com/gocolly/colly"
)

// @Author: Feng
// @Date: 2022/4/8 17:38

var (
	factory = Strategy{}
)

type Strategy struct {
	scraperMap map[string]Scraper
}

func GetStrategyInstance() Strategy {
	return factory
}

func (f Strategy) Register(name string, scraper Scraper) {
	f.scraperMap[name] = scraper
}

//Scraper colly封装
type Scraper interface {
	Init()
	ScrapeAndFlush()
}

//InitializeAndRegisterScraper 初始化并构造Scraper
func InitializeAndRegisterScraper(name string, cb func(*colly.Collector)) {
	// 初始化scraper
	s := InitialScraper(name, cb)
	// 注册scraper
	GetStrategyInstance().Register(name, s)
}
