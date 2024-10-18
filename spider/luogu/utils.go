package luogu

import (
	"fmt"
	"strconv"
)

func getPassAmountFromHtml(html string) int {
	s := string(html[39 : len(html)-7])
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("parse string to int error: %v", err)
	}
	return num
}

func getLuoguUrl(uid string) string {
	return fmt.Sprintf("https://www.luogu.com.cn/user/%s", uid)
}

func getLuoguUserUrl(name string) string {
	return fmt.Sprintf("https://www.luogu.com.cn/record/list?user=%v&page=1", name)
}

func getLuoguUserInfoHTMLSelector() string {
	return ".full-container"
	//return ".full-container .border.table .row .user"
}

func getLuoguPassAmountHTMLSelector() string {
	return ".stats.normal div:nth-child(4) span:last-child"
}

func getAmountPassKey(uid string) string {
	return "gxuicpc:luogu_problem_pass_amount_" + uid
}
