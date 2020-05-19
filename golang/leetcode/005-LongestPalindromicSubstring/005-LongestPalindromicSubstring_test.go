package longestpalindromicsubstring

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test1", args{s: ""}, ""},
		{"Test2", args{s: "a"}, ""},
		{"Test3", args{s: "aa"}, "aa"},
		{"Test4", args{s: "aab"}, "aa"},
		{"Test5", args{s: "baab"}, "baab"},
		{"Test6", args{s: "baaba"}, "baab"},
		{"Test7", args{s: "cbaaba"}, "baab"},
		{"Test8", args{s: "abaaba"}, "abaaba"},
		{"Test9", args{s: "babad"}, "bab"},
		{"Test10", args{s: "cbbd"}, "bb"},
		{"Test11", args{s: "abcdcbmnabcdcbaedf"}, "abcdcba"},
		{"Test12", args{s: "abcdefghhgfedcba"}, "abcdefghhgfedcba"},
		{"Test13", args{s: "abcdefghgfedcba"}, "abcdefghgfedcba"},
		{"Test14", args{s: "asabcdefghgfedcbamn"}, "abcdefghgfedcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}

			if got := longestPalindromeV2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
