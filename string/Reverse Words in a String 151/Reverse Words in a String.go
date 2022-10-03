package Reverse_Words_in_a_String_151

import "strings"

func reverseWords(s string) string {
	s := strings.Trim(s, " ")
	sli := strings.Split(s, " ")
	n := len(sli)
	ans := ""
	for i := len(sli); i >= 0; i-- {
		ans += sli[i]
	}
	return ans
}
