package atcoder

import (
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	submissionKey = "submission"
)

var (
	conScraper = scraper.NewScraper(
		scraper.WithCallback(conCallback),
		scraper.WithThreads(2),
	)
	contestId string
	userId    string
)

// submission 信息
type submission struct {
	userid string //用户名
	SMid   string //提交编号
	CTid   string //比赛编号
	task   string //题目序号
	score  int    //题目难度
}

//conCallback 处理比赛列表的回调函数
func conCallback(c *colly.Collector, res *scraper.Processor) {
	nowPage := 1
	// 遍历所有 contestPage
	c.OnHTML("div[class=\"col-lg-9 col-md-8\"]", func(element *colly.HTMLElement) {
		maxPage, err := strconv.Atoi(element.DOM.Find("div[class=\"text-center\"] ul li:last-child").First().Text())
		if err != nil {
			log.Errorf("Atcoder Page Error %v", err)
		}
		// 访问每个页面的contest
		element.ForEach("tbody tr", func(i int, element *colly.HTMLElement) {
			cLink := element.ChildAttr("td:nth-child(2) a", "href")
			contestId = strings.Split(cLink, "/")[2]

			// 递归进入用户比赛提交页面
			err = c.Visit(getAtCoderUrl(userId, contestId))
			if err != nil && err.Error() != "Not Found" {
				log.Errorf("Atcoder Page Error %v", err)
			}
		})
		//防止进入不存在的页面
		if nowPage < maxPage {
			nowPage = nowPage + 1
			err = c.Visit(getAtCoderPageUrl(nowPage))
			if err != nil && err.Error() != "Not Found" {
				log.Errorf("Atcoder Page Error %v", err)
			}
		}
	})

	//获取用户比赛提交页面信息
	c.OnHTML("table[class=\"table table-bordered table-striped small th-center\"] tbody tr", func(element *colly.HTMLElement) {
		//题目序号
		task := strings.Split(element.DOM.Find("td:nth-child(2)").First().Text(), "")[0]
		//题目难度
		score, errSc := strconv.Atoi(element.DOM.Find("td:nth-child(5)").First().Text())
		//提交编号
		SMid := element.ChildAttr("td:nth-child(10) a", "href")
		SMid = strings.Split(SMid, "/")[len(strings.Split(SMid, "/"))-1]

		if errSc != nil {
			log.Errorf("Submission Score Fetcher Error %v", errSc)
		}
		res.Set(submissionKey, submission{userId, SMid, contestId, task, score})
	})
}

//getAtCoderPageUrl 获取 userID
func getAtCoderPageUrl(page int) string {
	return "https://atcoder.jp/contests/archive?page=" + strconv.Itoa(page)
}

//getAtCoderUrl 获取用户提交页面链接
func getAtCoderUrl(atCoderId string, contestId string) string {
	return "https://atcoder.jp/contests/" + contestId + "/submissions?f.User=" + atCoderId + "&f.Status=AC"
}

//-------------------------------------------------------------------------------------------//
// 对外暴露函数
//-------------------------------------------------------------------------------------------//

//fetchConPage 抓取用户提交所有提交信息
func fetchConPage(uid string) ([]scraper.KV, error) {
	userId = uid
	//进入所有的比赛列表页面
	return conScraper.Scrape(getAtCoderPageUrl(1))
}
