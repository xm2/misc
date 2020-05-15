package longestpalindromicsubstring

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func isPalindrome(s string) bool {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-i-1] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	segmentSize := len(s)

	if segmentSize < 2 {
		return ""
	}

	for sz := segmentSize; sz > 1; sz-- {
		segments := len(s) - sz + 1
		for ss := 0; ss < segments; ss++ {
			segmentSlices := s[ss : ss+sz]
			if isPalindrome(segmentSlices) {
				return segmentSlices
			}
		}
	}

	return ""
}
