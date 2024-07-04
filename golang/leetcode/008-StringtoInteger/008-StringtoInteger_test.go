package stringtointeger

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test1", args{str: "42"}, 42},
		{"Test2", args{str: "   -42"}, -42},
		{"Test3", args{str: "4193 with words"}, 4193},
		{"Test4", args{str: "words and 987"}, 0},
		{"Test5", args{str: "-91283472332"}, -2147483648},
		{"Test6", args{str: "91283472332"}, 2147483647},
		{"Test7", args{str: "91283472332"}, 2147483647},
		{"Test8", args{str: "-2147483648"}, -2147483648},
		{"Test9", args{str: "2147483647"}, 2147483647},
		{"Test10", args{str: "-2147483647"}, -2147483647},
		{"Test11", args{str: "2147483646"}, 2147483646},
		{"Test12", args{str: "-2147483649"}, -2147483648},
		{"Test13", args{str: "2147483648"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
