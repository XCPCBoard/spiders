package luogu

import (
	"XCPCer_board/scraper"
	"context"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"time"
)

func getAmountPass(uid string) (scraper.KV, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var html string

	err := chromedp.Run(ctx,
		chromedp.Navigate(getLuoguUrl(uid)),
		chromedp.WaitVisible(getLuoguPassAmountHTMLSelector(), chromedp.ByQuery),
		chromedp.OuterHTML(getLuoguPassAmountHTMLSelector(), &html, chromedp.ByQuery),
	)
	if err != nil {
		log.Error(err)
		return scraper.KV{}, err
	}
	num := getPassAmountFromHtml(html)
	//关闭浏览器
	cancel()

	return scraper.KV{getAmountPassKey(uid), num}, nil
}

func GetProblemKvs(uid string) []scraper.KV {
	res := []scraper.KV{}
	pass, err := getAmountPass(uid)
	res = append(res, pass)
	if err != nil {
		log.Errorf("get amount pass problem error: %v", err)
	}

	return res
}
