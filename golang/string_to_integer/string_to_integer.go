package stringtointeger

import "strings"

func myAtoi(s string) int {
	var digits []byte = []byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
	}

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

	if s[0] == digits[0] {
		return 0
	}

	var i int = 1
	for j := range s {
		for d := range digits {
			if s[j] == digits[d] {
				i *= d
			}
		}
	}

	if negative {
		return -1
	}

	return 1
}
