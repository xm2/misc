package zigzagconversion

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test1", args{s: "PAYPALISHIRING", numRows: 1}, "PAYPALISHIRING"},
		{"Test2", args{s: "PAYPALISHIRING", numRows: 2}, "PYAIHRNAPLSIIG"},
		{"Test3", args{s: "PAYPALISHIRING", numRows: 3}, "PAHNAPLSIIGYIR"},
		{"Test4", args{s: "PAYPALISHIRING", numRows: 4}, "PINALSIGYAHRPI"},
		{"Test5", args{s: "PAYPALISHIRING", numRows: 5}, "PHASIYIRPLIGAN"},
		{"Test6", args{s: "PAYPALISHIRING", numRows: 6}, "PRAIIYHNPSGAIL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
