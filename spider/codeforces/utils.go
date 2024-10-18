package codeforces

import "fmt"

//---------------------------------------------------------------------//
// 常量
//---------------------------------------------------------------------//

const (
	// 个人rating
	ratingKey = "gxuicpc:codeforces_rating"
	// 个人历史最高rating
	maxRatingKey = "gxuicpc:codeforces_max_rating"
	//当前rating所对应的等级（红名、紫名...)
	rankingNameKey = "gxuicpc:codeforces_ranking_name"
	//最大rating所对应的等级（红名、紫名...)
	maxRankingNameKey = "gxuicpc:codeforces_max_ranking_name"
	// 个人总过题数
	problemPassAmountKey = "gxuicpc:codeforces_problem_pass"
	// 个人最后一月过题数
	lastMonthPassAmount = "gxuicpc:codeforces_last_month_problem_pass"

	// CF finder关键词
	// 个人总过题数
	problemPassKeyWord = "all"
	//个人最后一月过题数
	lastMonthPassKeyWord = "month"
)

// ---------------------------------------------------------------------//
// 共用函数
// ---------------------------------------------------------------------//
// 使用镜像网站，躲避官网人机校验
func getPersonPageUrl(uid string) string {
	return "https://codeforc.es/profile/" + uid
}

func getUserInfoUrl(uid string) string {
	return "https://codeforc.es/api/user.info?handles=" + uid
}

func GetRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", ratingKey, uid)
}

func GetMaxRatingKey(uid string) string {
	return fmt.Sprintf("%v_%v", maxRatingKey, uid)
}

func GetRankingNameKey(uid string) string {
	return fmt.Sprintf("%v_%v", rankingNameKey, uid)
}

func GetMaxRankingNameKey(uid string) string {
	return fmt.Sprintf("%v_%v", maxRankingNameKey, uid)
}
