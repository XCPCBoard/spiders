package nowcoder

import "fmt"

// @Author: Feng
// @Date: 2022/5/16 17:48

//---------------------------------------------------------------------//
// 常量
//---------------------------------------------------------------------//

// 牛客finder存储Key
const (
	// 个人主页selector关键字
	ratingKeyWord        = "Rating"
	ratingRankingKeyWord = "Rating排名"
	contestAmountKeyWord = "次比赛"
	// 个人练习selector关键字
	passAmountKeyWord = "题已通过"
)

//---------------------------------------------------------------------//
// 共用函数
//---------------------------------------------------------------------//

//getContestProfileUrl 获取牛客竞赛区个人主页URL
func getContestProfileUrl(nowCoderId string) string {
	return "https://ac.nowcoder.com/acm/contest/profile/" + nowCoderId
}

//getContestPracticeUrl 获取牛客竞赛区个人练习URL
func getContestPracticeUrl(nowCoderId string) string {
	return getContestProfileUrl(nowCoderId) + "/practice-coding"
}

//getNowCoderContestBaseFindRule 获取牛客竞赛区基础的
func getNowCoderContestBaseFindRule(keyWord string) string {
	return fmt.Sprintf(".my-state-item:contains(%v) .state-num", keyWord)
}
