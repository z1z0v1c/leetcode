package stringtointeger

import "strings"

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	var negative bool
	if s[0] == '-' {
		negative = true
	}

	s = strings.Trim(s, "+-")
	if len(s) == 0 {
		return 0
	}

	if negative {
		return -1
	}

	return 1
}
