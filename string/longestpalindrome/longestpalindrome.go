package longestpalindrome

func LongestPalindromicSubstring(str string) string {
	longest := ""
	for i := range str {
		for j := i; j < len(str); j++ {
			substr := str[i : j+1]
			if len(substr) > len(longest) && isPalindrome(substr) {
				longest = substr
			}
		}
	}
	return longest
}

func isPalindrome(str string) bool {
	for i := range str {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
