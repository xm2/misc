package regularexpressionmatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Test1", args{s: "aa", p: "a"}, false},
		{"Test2", args{s: "aa", p: "a*"}, true},
		{"Test3", args{s: "ab", p: ".*"}, true},
		{"Test4", args{s: "aab", p: "c*a*b"}, true},
		{"Test5", args{s: "mississippi", p: "mis*is*p*."}, false},
		{"Test6", args{s: "mississippi", p: "mis*is*ip*."}, true},
		{"Test7", args{s: "mississippi", p: "mis.*p"}, false},
		{"Test8", args{s: "mississippi", p: "mis.*p*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
