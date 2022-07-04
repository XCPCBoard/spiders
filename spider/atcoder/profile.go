package atcoder

import (
	"XCPCer_board/scraper"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

//---------------------------------------------------------------------//
// atCoder个人信息 //
//---------------------------------------------------------------------//
//  Key

const (
	//key

	RatingKey     = "atc_rating"
	contestSumKey = "atc_contest_sum"
	rankKey       = "atc_rank"

	//keyword
)

var (
	mainScraper = scraper.NewScraper(
		scraper.WithCallback(mainCallback),
		scraper.WithThreads(2),
	)
)

//mainCallback 处理个人主页的回调函数
func mainCallback(c *colly.Collector, res *scraper.Processor) {
	c.OnHTML("table[class=\"dl-table mt-2\"] tbody",
		func(element *colly.HTMLElement) {
			// 获取rating
			retRating := element.DOM.Find(fmt.Sprintf("tr:nth-child(2) span:first-child")).First().Text()
			if num, err := strconv.Atoi(retRating); err == nil {
				res.Set(RatingKey, num)
			}
			// 获取Rank
			retRank := strings.Split(element.DOM.Find(fmt.Sprintf("tr:nth-child(1) td")).First().Text(), "th")[0]
			if num, err := strconv.Atoi(retRank); err == nil {
				res.Set(rankKey, num)
			}
			// 获取rating比赛场数
			retConSum := element.DOM.Find(fmt.Sprintf("tr:nth-child(4) td")).First().Text()
			if num, err := strconv.Atoi(retConSum); err == nil {
				res.Set(contestSumKey, num)
			}
		},
	)
}

//getAtCoderBaseUrl 获取个人主页URL
func getAtCoderBaseUrl(atCoderId string) string {
	return "https://atcoder.jp/users/" + atCoderId
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchMainPage 抓取个人主页页面所有
func fetchMainPage(uid string) ([]scraper.KV, error) {
	return mainScraper.Scrape(getAtCoderBaseUrl(uid))
}
