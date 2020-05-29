package longestcommonprefix

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) < 2 {
		return ""
	}

	ret := ""
	for i := 0; i < len(strs[0]); i++ {
		cs := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || cs != strs[j][i] {
				return ret
			}
		}
		ret += string(cs)
	}

	return ret

}
