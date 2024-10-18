package vjudge

import "fmt"

func getVjudgeAmountPass(uid string) string {
	return fmt.Sprintf("gxuicpc:vjudge_problem_amount_pass_%v", uid)
}

func getProfileUrl(uid string) string {
	return fmt.Sprintf("https://vjudge.net/user/%v", uid)
}
