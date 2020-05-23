package containerwithmostwater

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test1", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"Test2", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 8}}, 64},
		{"Test3", args{height: []int{1, 8, 6, 2, 5, 4, 8, 3}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
			if got := maxAreaV2(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
