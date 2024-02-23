package longestpalindromesubstring

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	lp := string(s[0])

	for i := 0; i < len(s); i++ {
		sub := string(s[i])
		for j := i + 1; j < len(s); j++ {
			sub += string(s[j])
			if isPalindrome(&sub) {
				if len(sub) > len(lp) {
					lp = sub
				}
			}
		}
	}

	return lp
}

func isPalindrome(s *string) bool {
	for i := 0; i < len(*s)/2; i++ {
		if (*s)[i] != (*s)[len(*s)-i-1] {
			return false
		}
	}
	return true
}
