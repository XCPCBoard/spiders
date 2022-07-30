package codeforces

//---------------------------------------------------------------------//
// 常量
//---------------------------------------------------------------------//

const (
	// CF finder关键词
	// 个人总过题数
	problemPassKeyWord = "all"
	//个人最后一月过题数
	lastMonthPassKeyWord = "month"
)

//---------------------------------------------------------------------//
// 共用函数
//---------------------------------------------------------------------//

func getPersonPageUrl(uid string) string {
	return "https://codeforces.com/profile/" + uid
}

func getUserInfoUrl(uid string) string {
	return "https://codeforces.com/api/user.info?handles=" + uid
}
