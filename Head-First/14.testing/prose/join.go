package prose

import "strings"

//JoinWithCommas accepts a slice of strings to join, return the resulted phrase
func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		res := strings.Join(phrases[:len(phrases)-1], ", ")
		res += ", and "
		res += phrases[len(phrases)-1]
		return res
	}
}
