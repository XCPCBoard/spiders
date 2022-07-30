package vjudge

//KeyWordListOfUser 用户keyWord常量列表
var KeyWordListOfUser = []string{
	last24HoursNumber, last7DaysNumber, last30DaysNumber, totalNumber,
}

//---------------------------------------------------------------------//
// 部分共用函数 //
//---------------------------------------------------------------------//

func getPersonPage(uid string) string {
	return "https://vjudge.net/user/" + uid
}
