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

/**
time complexity O(n**3)
*/
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

/**
time complexity O(n**2)
*/
func longestPalindromeV2(s string) string {
	maxLength := 0
	centerIndex := 0

	maxLenPalindromeAround := func(s string, i int, j int) int {
		for {
			if i >= 0 && j < len(s) && s[i] == s[j] {
				i--
				j++
			} else {
				break
			}
		}
		return j - i - 1
	}

	for i := 0; i < len(s); i++ {
		len1 := maxLenPalindromeAround(s, i, i)
		len2 := maxLenPalindromeAround(s, i, i+1)

		if len1 < len2 {
			// len1 carry the bigger value to simplify the blow compare logic
			len1 = len2
		}

		if len1 > maxLength {
			maxLength = len1
			centerIndex = i
		}
	}

	if maxLength < 2 {
		// exclude single letter
		return ""
	}

	return s[centerIndex-(maxLength-1)/2 : centerIndex+maxLength/2+1]

}
