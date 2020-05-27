package integertoroman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test1", args{num: 3}, "III"},
		{"Test2", args{num: 4}, "IV"},
		{"Test3", args{num: 9}, "IX"},
		{"Test4", args{num: 58}, "LVIII"},
		{"Test5", args{num: 1994}, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
