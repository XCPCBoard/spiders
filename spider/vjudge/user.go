package vjudge

import (
	"XCPCer_board/model"
	"XCPCer_board/scraper"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strconv"
)

//---------------------------------------------------------------------//
// 获取int形式的信息 //
//---------------------------------------------------------------------//

var (
	userScraper = scraper.NewScraper(
		profileCallback,
	)
)

func profileCallback(c *colly.Collector) {
	c.OnHTML(".container div:nth-child(2) .table.table-reflow.problem-solve ",
		func(e *colly.HTMLElement) {
			uid := e.Request.Ctx.Get("uid")
			if uid == "" {
				log.Errorf("%v", model.UidError)
				return
			}
			s := e.DOM.Find("a[title=\"Overall solved\"]").First().Text()
			//log.Printf(s)
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Errorf("parse string to int error: %v", err)
			} else {
				e.Request.Ctx.Put(getVjudgeAmountPass(uid), num)
			}
		},
	)
}

// //////////////////////对外暴露函数////////////////////////
func fetchProfile(uid string) ([]scraper.KV, error) {
	ctx := colly.NewContext()
	ctx.Put("uid", uid)
	//err := mainScraper.C.Request("GET", getContestProfileUrl(uid), nil, ctx, nil)

	err := userScraper.C.Request("GET", getProfileUrl(uid), nil, ctx, nil)
	if err != nil {
		log.Errorf("vjudge fetch error: %v", err)
		return nil, err
	}
	kvs := scraper.Parse(ctx, map[string]struct{}{
		"uid": {},
	})
	//log.Printf("kvs: %v", kvs)
	return kvs, nil
}
