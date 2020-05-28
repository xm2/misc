package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test1", args{s: "III"}, 3},
		{"Test2", args{s: "IV"}, 4},
		{"Test3", args{s: "IX"}, 9},
		{"Test3", args{s: "LVIII"}, 58},
		{"Test3", args{s: "MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
