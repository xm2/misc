package longestsubstring

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	start, end, longest := 0, 0, 0
	cmap := map[rune]int{}

	for i, v := range s {
		if _, ok := cmap[v]; ok {
			if end-start > longest {
				longest = end - start
			}
			start = end
			cmap = map[rune]int{v: i}
			end++
		} else {
			end++
			cmap[v] = i
		}
	}

	if end-start > longest {
		longest = end - start
	}

	return longest
}
