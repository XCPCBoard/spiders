package luogu

import (
	"context"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"time"
)

func getUserId(name string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var html string

	err := chromedp.Run(ctx,
		chromedp.Navigate(getLuoguUserUrl(name)),
		chromedp.WaitVisible(getLuoguUserInfoHTMLSelector(), chromedp.ByQuery),
		chromedp.OuterHTML(getLuoguUserInfoHTMLSelector(), &html, chromedp.ByQuery),
	)
	if err != nil {
		log.Error(err)
		return "", err
	}
	//uid := getPassAmountFromHtml(html)
	log.Infof("Html content: %v", html)
	return "uid", nil
}
