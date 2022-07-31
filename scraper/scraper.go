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
	Scrape(ctx *colly.Context)
}
