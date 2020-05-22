package palindromenumber

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Test1", args{x: 121}, true},
		{"Test2", args{x: -121}, false},
		{"Test3", args{x: 10}, false},
		{"Test4", args{x: 11}, true},
		{"Test5", args{x: 1111}, true},
		{"Test6", args{x: 11211}, true},
		{"Test7", args{x: 12121}, true},
		{"Test8", args{x: 1}, true},
		{"Test9", args{x: 110}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
